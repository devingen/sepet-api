package s3fs

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	core "github.com/devingen/api-core"
	"github.com/devingen/sepet-api/model"
	"strings"
)

// GetFileList implements ISepetService interface
func (s3Service S3Service) GetFileList(ctx context.Context, bucket *model.Bucket, bucketVersion, path string, fetchOnlyDirectChildren bool) ([]model.File, error) {

	// all files under the folder
	fileList, err := fetchFileList(s3Service, core.StringValue(bucket.Folder), bucketVersion, path)
	if err != nil {
		return nil, err
	}

	if fetchOnlyDirectChildren {
		// filter the files to get only the direct children
		return getContentsOfDirectory(fileList, path), nil
	}
	return fileList, nil
}

func fetchFileList(s3Service S3Service, bucketName, version, path string) ([]model.File, error) {
	sess := session.New(s3Service.Config)
	s3Client := s3.New(sess, s3Service.Config)

	prefix := bucketName
	pathPrefix := bucketName
	if version != "" {
		prefix += "/" + version
		pathPrefix += "/" + version
	}
	if path != "" {
		prefix += "/" + path
	}
	//prefixLength := len(prefix) + 1
	pathPrefixLength := len(pathPrefix) + 1

	files := make([]model.File, 0)
	nextMarker := ""
	for {
		input := &s3.ListObjectsInput{
			Bucket: aws.String(s3Service.Bucket),
			Prefix: aws.String(prefix),
		}

		if nextMarker != "" {
			input.Marker = aws.String(nextMarker)
		}

		result, err := s3Client.ListObjects(input)
		if err != nil {
			return nil, err
		}
		resultLength := len(result.Contents)

		for i := 0; i < resultLength; i++ {

			filePath := *result.Contents[i].Key
			name := filePath[strings.LastIndex(filePath, "/")+1:]
			path := filePath[pathPrefixLength:]
			files = append(files, model.File{
				Name:         name,
				Path:         path,
				Size:         *result.Contents[i].Size,
				LastModified: *result.Contents[i].LastModified,
			})
		}

		if !*result.IsTruncated {
			// break if there are no more files to fetch
			break
		}

		if result.NextMarker != nil {
			// use the next marker returned in the response to skip already fetched files
			nextMarker = *result.NextMarker
		} else {
			// use the last file path as next marker
			nextMarker = *result.Contents[resultLength-1].Key
		}
	}
	return files, nil
}

func getContentsOfDirectory(files []model.File, path string) []model.File {
	pathLength := len(path) + 1
	contentsSet := make(map[string]model.File)
	for _, file := range files {
		namePath := file.Path
		if path != "" {
			// clear the prefix if there is path
			if index := strings.Index(file.Path, path+"/"); index == 0 {
				namePath = file.Path[index+pathLength:]
			} else {
				// skip the file if it doesn't match the prefix
				continue
			}
		}

		if index := strings.Index(namePath, "/"); index != -1 {
			// this is a file in one of the sub directories. clear the name and path to have the directory name only
			subFilePath := path + "/" + namePath[:index]
			if path == "" {
				subFilePath = namePath[:index]
			}
			contentsSet[namePath[:index+1]] = model.File{
				Name:         namePath[:index],
				Path:         subFilePath,
				Size:         file.Size,
				LastModified: file.LastModified,
			}
		} else {
			contentsSet[namePath] = file
		}
	}
	contents := make([]model.File, 0)
	for filePath, file := range contentsSet {
		if filePath != "" {
			contents = append(contents, file)
		}
	}
	return contents
}
