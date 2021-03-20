package s3fs

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/devingen/sepet-api/model"
)

// DeleteFile implements ISepetService interface
func (s3Service S3Service) DeleteFile(ctx context.Context, bucket *model.Bucket, bucketVersion, path string) error {

	sess := session.New(s3Service.Config)
	s3Client := s3.New(sess, s3Service.Config)

	// it's a folder if the path ends with /
	isFolder := len(path) > 0 && path[len(path)-1:] == "/"
	if isFolder {
		files, gfErr := s3Service.GetFileList(ctx, bucket, bucketVersion, path[:len(path)-1], false)
		if gfErr != nil {
			return gfErr
		}

		for _, filePath := range files {
			_, dfErr := s3Client.DeleteObject(&s3.DeleteObjectInput{
				Bucket: aws.String(s3Service.Bucket),
				Key:    aws.String(GetFilePath(bucket.Folder, bucketVersion, path+filePath)),
			})
			if dfErr != nil {
				return dfErr
			}
		}
	}

	_, err := s3Client.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(s3Service.Bucket),
		Key:    aws.String(GetFilePath(bucket.Folder, bucketVersion, path)),
	})
	return err
}
