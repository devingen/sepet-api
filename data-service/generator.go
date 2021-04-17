package ds

import (
	"github.com/devingen/api-core/database"
	mongo_data_service "github.com/devingen/sepet-api/data-service/mongo-data-service"
)

// New generates new MongoDataService
func New(databaseName string, database *database.Database) ISepetDataService {
	return mongo_data_service.MongoDataService{
		DatabaseName: databaseName,
		Database:     database,
	}
}
