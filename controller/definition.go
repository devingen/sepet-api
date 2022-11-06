package controller

import (
	"context"
	core "github.com/devingen/api-core"
)

// IServiceController defines the functionality of the service controller
type IServiceController interface {
	CreateBucket(ctx context.Context, req core.Request) (*core.Response, error)
	DeleteBucket(ctx context.Context, req core.Request) (*core.Response, error)
	DeleteFile(ctx context.Context, req core.Request) (*core.Response, error)
	GetBucket(ctx context.Context, req core.Request) (*core.Response, error)
	GetBucketVersionList(ctx context.Context, req core.Request) (*core.Response, error)
	GetBuckets(ctx context.Context, req core.Request) (*core.Response, error)
	GetFileList(ctx context.Context, req core.Request) (*core.Response, error)
	UpdateBucket(ctx context.Context, req core.Request) (*core.Response, error)
	UploadFile(ctx context.Context, req core.Request) (*core.Response, error)

	// This function is created because AWS cannot use
	// multiple handlers to one endpoint which catches all the requests.
	// See "Catch-all Path Variables" in
	// https://aws.amazon.com/blogs/aws/api-gateway-update-new-features-simplify-api-development/
	GetFileListOrDeleteFile(ctx context.Context, req core.Request) (*core.Response, error)
}
