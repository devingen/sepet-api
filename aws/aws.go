package aws

import (
	"context"
	core "github.com/devingen/api-core"
	"github.com/devingen/api-core/database"
	"github.com/devingen/api-core/wrapper"
	"github.com/devingen/sepet-api/config"
	"github.com/devingen/sepet-api/controller"
	service_controller "github.com/devingen/sepet-api/controller/service-controller"
	data_service "github.com/devingen/sepet-api/data-service"
	file_service "github.com/devingen/sepet-api/file-service"
	"github.com/kelseyhightower/envconfig"
	"log"
)

var db *database.Database

// InitDeps creates the dependencies for the AWS Lambda functions.
func InitDeps() (*service_controller.ServiceController, func(f core.Controller) wrapper.AWSLambdaHandler) {
	var appConfig config.App
	err := envconfig.Process("sepet", &appConfig)
	if err != nil {
		log.Fatal(err.Error())
	}

	dataService := data_service.New(appConfig.Mongo.Database, getDatabase(appConfig))
	s3Service := file_service.New(appConfig.S3, appConfig.CDNDomain, appConfig.CDNProtocol)
	serviceController := controller.New(dataService, s3Service)

	wrap := generateWrapper(appConfig)

	return serviceController, wrap
}

func getDatabase(appConfig config.App) *database.Database {
	if db == nil {
		var err error
		db, err = database.New(appConfig.Mongo.URI)
		if err != nil {
			log.Fatalf("Database connection failed %s", err.Error())
		}
	} else if !db.IsConnected() {
		err := db.ConnectWithEnvironment()
		if err != nil {
			log.Fatalf("Database connection failed %s", err.Error())
		}
	} else {
		log.Println("Database connection exists")
	}
	return db
}

func generateWrapper(appConfig config.App) func(f core.Controller) wrapper.AWSLambdaHandler {
	return func(f core.Controller) wrapper.AWSLambdaHandler {
		ctx := context.Background()

		// add logger
		withLogger := wrapper.WithLogger(appConfig.LogLevel, f)

		// convert to HTTP handler
		handler := wrapper.WithLambdaHandler(ctx, withLogger)
		return handler
	}
}
