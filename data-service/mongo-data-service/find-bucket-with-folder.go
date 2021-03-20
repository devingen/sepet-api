package mongods

import (
	"context"
	"github.com/devingen/sepet-api/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// FindBucketWithFolder implements ISepetDataService interface
func (service MongoDataService) FindBucketWithFolder(ctx context.Context, folder string) (*model.Bucket, error) {
	result := model.Bucket{}
	err := service.Database.FindOne(ctx, service.DatabaseName, CollectionBuckets, bson.M{"folder": folder}, &result)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	return &result, err
}
