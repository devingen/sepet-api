package mongods

import (
	"context"
	"github.com/devingen/sepet-api/model"
	"go.mongodb.org/mongo-driver/mongo"
)

// FindBucketWithID implements ISepetDataService interface
func (service MongoDataService) FindBucketWithID(ctx context.Context, id string) (*model.Bucket, error) {
	result := model.Bucket{}
	err := service.Database.Get(ctx, service.DatabaseName, CollectionBuckets, id, &result)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	return &result, err
}
