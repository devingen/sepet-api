package mongods

import (
	"context"
	"github.com/devingen/sepet-api/model"
)

// CreateBucket implements ISepetDataService interface
func (service MongoDataService) CreateBucket(ctx context.Context, item *model.Bucket) (*model.Bucket, error) {

	item.AddCreationFields()

	id, err := service.Database.Create(ctx, service.DatabaseName, CollectionBuckets, item)
	if err != nil {
		return nil, err
	}

	item.ID = *id
	return item, nil
}
