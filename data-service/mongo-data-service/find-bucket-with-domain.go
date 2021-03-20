package mongods

import (
	"context"
	"github.com/devingen/sepet-api/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// FindBucketWithDomain implements ISepetDataService interface
func (service MongoDataService) FindBucketWithDomain(ctx context.Context, domain string) (*model.Bucket, error) {
	result := model.Bucket{}
	err := service.Database.FindOne(ctx, service.DatabaseName, CollectionBuckets, bson.M{"domain": domain}, &result)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	return &result, err
}
