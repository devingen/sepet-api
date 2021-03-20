package mongods

import "github.com/devingen/api-core/database"

const (
	// CollectionBuckets is the MongoDB collection of the buckets
	CollectionBuckets = "sepet-buckets"
)

// MongoDataService implements ISepetDataService interface with MongoDB connection
type MongoDataService struct {
	DatabaseName string
	Database     *database.Database
}
