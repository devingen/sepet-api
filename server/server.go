package server

import (
	"context"
	core "github.com/devingen/api-core"
	"github.com/devingen/api-core/database"
	"github.com/devingen/api-core/server"
	"github.com/devingen/api-core/wrapper"
	"github.com/devingen/sepet-api/config"
	"github.com/devingen/sepet-api/controller"
	data_service "github.com/devingen/sepet-api/data-service"
	file_service "github.com/devingen/sepet-api/file-service"
	customvalidator "github.com/devingen/sepet-api/validator"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"net/http"
)

// New creates a new HTTP server
func New(appConfig config.App, db *database.Database) *http.Server {

	validate := validator.New()
	validate.RegisterValidation("bucket-domain", customvalidator.ValidateBucketDomain)
	core.SetValidator(validate)

	srv := &http.Server{Addr: ":" + appConfig.Port}

	dataService := data_service.New(appConfig.Mongo.Database, db)
	s3Service := file_service.New(appConfig.S3, appConfig.CDNDomain, appConfig.CDNProtocol)
	serviceController := controller.New(dataService, s3Service)

	wrap := generateWrapper(appConfig)

	router := mux.NewRouter()
	router.HandleFunc("/buckets", wrap(serviceController.GetBuckets)).Methods(http.MethodGet)
	router.HandleFunc("/buckets", wrap(serviceController.CreateBucket)).Methods(http.MethodPost)
	router.HandleFunc("/buckets/{id}", wrap(serviceController.GetBucket)).Methods(http.MethodGet)
	router.HandleFunc("/buckets/{id}", wrap(serviceController.UpdateBucket)).Methods(http.MethodPut)
	router.HandleFunc("/buckets/{id}", wrap(serviceController.DeleteBucket)).Methods(http.MethodDelete)
	router.HandleFunc("/{domain}", wrap(serviceController.UploadFile)).Methods(http.MethodPost)
	router.PathPrefix("/").Handler(http.HandlerFunc(wrap(serviceController.DeleteFile))).Methods(http.MethodDelete)
	router.PathPrefix("/").Handler(http.HandlerFunc(wrap(serviceController.GetFileList))).Methods(http.MethodGet)

	http.Handle("/", &server.CORSRouterDecorator{R: router})
	return srv
}

func generateWrapper(appConfig config.App) func(f core.Controller) func(http.ResponseWriter, *http.Request) {
	return func(f core.Controller) func(http.ResponseWriter, *http.Request) {
		ctx := context.Background()

		// add logger
		withLogger := wrapper.WithLogger(appConfig.LogLevel, f)

		// convert to HTTP handler
		handler := wrapper.WithHTTPHandler(ctx, withLogger)
		return handler
	}
}
