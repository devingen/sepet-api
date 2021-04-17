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
func (s3Service S3Service) GetFileList(ctx context.Context, bucket *model.Bucket, bucketVersion, path string, fetchOnlyDirectChildren bool) ([]string, error) {

	// all files under the folder
	fileList, err := fetchFileList(s3Service, core.StringValue(bucket.Folder), bucketVersion, path)
	if err != nil {
		return nil, err
	}

	if fetchOnlyDirectChildren {
		// filter the files to get only the direct children
		return getContentsOfDirectory(fileList, ""), nil
	}
	return fileList, nil
}

func fetchFileList(s3Service S3Service, bucketName, version, path string) ([]string, error) {
	sess := session.New(s3Service.Config)
	s3Client := s3.New(sess, s3Service.Config)

	prefix := bucketName
	if version != "" {
		prefix += "/" + version
	}
	if path != "" {
		prefix += "/" + path
	}
	prefixLength := len(prefix) + 1

	files := make([]string, 0)
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
			files = append(files, filePath[prefixLength:])
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

func getContentsOfDirectory(files []string, path string) []string {
	pathLength := len(path) + 1
	contentsSet := make(map[string]bool)
	for _, file := range files {
		filePath := file
		if path != "" {
			// clear the path prefix if there is path
			if index := strings.Index(filePath, path+"/"); index == 0 {
				filePath = filePath[index+pathLength:]
			} else {
				// skip the file if it doesn't match the prefix
				continue
			}
		}

		if index := strings.Index(filePath, "/"); index != -1 {
			contentsSet[filePath[:index+1]] = true
			//contentsSet[filePath] = true
		} else {
			contentsSet[filePath] = true
		}
	}
	contents := make([]string, 0)
	for i := range contentsSet {
		if i != "" {
			contents = append(contents, i)
		}
	}
	return contents
}
