package mongods

import (
	"context"
	"github.com/devingen/sepet-api/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// FindBuckets implements ISepetDataService interface
func (service MongoDataService) FindBuckets(ctx context.Context, status model.BucketStatus) ([]*model.Bucket, error) {
	query := bson.M{}
	if status != "" {
		query["status"] = status
	}

	result := make([]*model.Bucket, 0)
	err := service.Database.Find(ctx, service.DatabaseName, CollectionBuckets, query, 0, func(cur *mongo.Cursor) error {
		var data model.Bucket
		err := cur.Decode(&data)
		if err != nil {
			return err
		}
		result = append(result, &data)
		return nil
	})
	return result, err
}
