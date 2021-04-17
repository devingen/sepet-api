package s3fs

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	core "github.com/devingen/api-core"
	"github.com/devingen/sepet-api/model"
)

// DeleteEntireBucket implements ISepetService interface
func (s3Service S3Service) DeleteEntireBucket(ctx context.Context, bucket *model.Bucket) error {

	sess := session.New(s3Service.Config)
	s3Client := s3.New(sess, s3Service.Config)

	files, gfErr := s3Service.GetFileList(ctx, bucket, "", "", false)
	if gfErr != nil {
		return gfErr
	}

	for _, filePath := range files {
		_, dfErr := s3Client.DeleteObject(&s3.DeleteObjectInput{
			Bucket: aws.String(s3Service.Bucket),
			Key:    aws.String(core.StringValue(bucket.Folder) + "/" + filePath),
		})
		if dfErr != nil {
			return dfErr
		}
	}
	return nil
}
