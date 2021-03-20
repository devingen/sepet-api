package mongods

import (
	"context"
	"github.com/devingen/sepet-api/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateBucket implements ISepetDataService interface
func (service MongoDataService) CreateBucket(ctx context.Context, item *model.Bucket) (*model.Bucket, error) {
	collection, err := service.Database.ConnectToCollection(service.DatabaseName, CollectionBuckets)
	if err != nil {
		return nil, err
	}

	item.AddCreationFields()

	result, err := collection.InsertOne(ctx, item)
	if err != nil {
		return nil, err
	}

	item.ID = result.InsertedID.(primitive.ObjectID)
	return item, nil
}
