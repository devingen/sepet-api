package mongods

import (
	"context"
	"github.com/devingen/sepet-api/model"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

// UpdateBucket implements ISepetDataService interface
func (service MongoDataService) UpdateBucket(ctx context.Context, item *model.Bucket) (*time.Time, int, error) {

	item.PrepareUpdateFields()

	var result = &model.Bucket{}
	err := service.Database.Update(ctx, service.DatabaseName, CollectionBuckets, item.ID, &result, bson.M{
		"$set": item,
		"$inc": bson.M{"_revision": 1},
	})
	if err != nil {
		return nil, 0, err
	}

	return result.UpdatedAt, result.Revision + 1, nil
}
