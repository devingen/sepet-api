package mongods

import (
	"context"
	"github.com/devingen/api-core/database"
	"github.com/devingen/sepet-api/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// FindBuckets implements ISepetDataService interface
func (service MongoDataService) FindBuckets(ctx context.Context, query bson.M) ([]*model.Bucket, error) {

	result := make([]*model.Bucket, 0)
	err := service.Database.Find(ctx, service.DatabaseName, CollectionBuckets, query, database.FindOptions{}, func(cur *mongo.Cursor) error {
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
