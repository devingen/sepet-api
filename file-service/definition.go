package fs

import (
	"context"
	"github.com/devingen/sepet-api/model"
	"mime/multipart"
)

// ISepetService defines the functionality of the file server
type ISepetService interface {
	GetFileList(ctx context.Context, bucket *model.Bucket, bucketVersion, path string, fetchOnlyDirectChildren bool) ([]string, error)
	DeleteFile(ctx context.Context, bucket *model.Bucket, bucketVersion, path string) error
	UploadFile(ctx context.Context, bucket *model.Bucket, bucketVersion string, files map[string]multipart.File) ([]string, error)
	DeleteEntireBucket(ctx context.Context, bucket *model.Bucket) error
}
