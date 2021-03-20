package mongods

import (
	"context"
)

// DeleteBucket implements ISepetDataService interface
func (service MongoDataService) DeleteBucket(ctx context.Context, id string) error {
	_, err := service.Database.Delete(ctx, service.DatabaseName, CollectionBuckets, id)
	return err
}
