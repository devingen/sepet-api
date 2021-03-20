package controller

import (
	"context"
	core "github.com/devingen/api-core"
)

// IServiceController defines the functionality of the service controller
type IServiceController interface {
	CreateBucket(ctx context.Context, req core.Request) (interface{}, int, error)
	DeleteBucket(ctx context.Context, req core.Request) (interface{}, int, error)
	DeleteFile(ctx context.Context, req core.Request) (interface{}, int, error)
	GetBucket(ctx context.Context, req core.Request) (interface{}, int, error)
	GetBuckets(ctx context.Context, req core.Request) (interface{}, int, error)
	GetFileList(ctx context.Context, req core.Request) (interface{}, int, error)
	UpdateBucket(ctx context.Context, req core.Request) (interface{}, int, error)
	UploadFile(ctx context.Context, req core.Request) (interface{}, int, error)

	// This function is created because AWS cannot use
	// multiple handlers to one endpoint which catches all the requests.
	// See "Catch-all Path Variables" in
	// https://aws.amazon.com/blogs/aws/api-gateway-update-new-features-simplify-api-development/
	GetFileListOrDeleteFile(ctx context.Context, req core.Request) (interface{}, int, error)
}
