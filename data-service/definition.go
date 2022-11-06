package ds

import (
	"context"
	"github.com/devingen/sepet-api/model"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

// ISepetDataService defines the functionality of the data service
type ISepetDataService interface {
	CreateBucket(ctx context.Context, item *model.Bucket) (*model.Bucket, error)
	UpdateBucket(ctx context.Context, item *model.Bucket) (*time.Time, int, error)
	DeleteBucket(ctx context.Context, id string) error
	FindBuckets(ctx context.Context, query bson.M) ([]*model.Bucket, error)
	FindBucketWithID(ctx context.Context, id string) (*model.Bucket, error)
	FindBucketWithDomain(ctx context.Context, domain string) (*model.Bucket, error)
	FindBucketWithFolder(ctx context.Context, folder string) (*model.Bucket, error)
}
