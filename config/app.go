package config

// App defines the environment variable configuration for the whole app
type App struct {
	// Port is the port of the HTTP server.
	Port string `envconfig:"api_port" default:"8080"`

	// CDNDomain is the domain to be used in the generated CDN url.
	// The bucket domain will be added as sub domain and the CDN protocol will be prepended.
	// Formula is CDN_PROTOCOL://BUCKET_DOMAIN.CDN_DOMAIN
	// Result  is https://acme.sepet.devingen.io
	CDNDomain string `envconfig:"cdn_domain" default:"sepet.devingen.io"`

	// CDNProtocol is the protocol to be used in the generated CDN url. Should be either http or https.
	CDNProtocol string `envconfig:"cdn_protocol" default:"https"`

	// LogLevel defines the log level.
	LogLevel string `envconfig:"api_log_level" default:"info"`

	// Mongo is the configuration of the MongoDB server.
	Mongo Mongo `envconfig:"mongo"`

	// S3 is the configuration of the S3 server.
	S3 S3 `envconfig:"s3"`
}

// Mongo defines the environment variable configuration for MongoDB
type Mongo struct {
	// URI is the MongoDB server URI.
	URI string `envconfig:"uri" default:"mongodb://localhost:27017"`

	// Database is the MongoDB database name to connect to.
	Database string `envconfig:"database" default:"sepet"`
}

// S3 defines the environment variable configuration for AWS S3 or MinIO
type S3 struct {
	// Endpoint is the URL of the file server to connect to. If empty, the connection is made to the AWS S3 servers.
	// Used to connect a local MinIO server for development and integration tests.
	Endpoint string `envconfig:"endpoint"`

	// AccessKeyID is the AWS access key ID.
	AccessKeyID string `envconfig:"access_key_id" required:"true"`

	// AccessKey is the AWS access key.
	AccessKey string `envconfig:"secret_access_key" required:"true"`

	// Region is the AWS region.
	Region string `envconfig:"region" required:"true"`

	// Bucket is the bucket to connect.
	Bucket string `envconfig:"bucket" default:"sepet"`
}
