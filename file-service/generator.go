package fs

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/devingen/sepet-api/config"
	s3_service "github.com/devingen/sepet-api/file-service/s3-file-service"
)

// New generates new S3Service
func New(envConfig config.S3, CDNDomain, CDNProtocol string) *s3_service.S3Service {

	if envConfig.Endpoint != "" {
		// configure to use MinIO Server
		return &s3_service.S3Service{
			Bucket:      envConfig.Bucket,
			CDNDomain:   CDNDomain,
			CDNProtocol: CDNProtocol,
			Config: &aws.Config{
				Credentials:      credentials.NewStaticCredentials(envConfig.AccessKeyID, envConfig.AccessKey, ""),
				Endpoint:         aws.String(envConfig.Endpoint),
				Region:           aws.String(envConfig.Region),
				DisableSSL:       aws.Bool(true),
				S3ForcePathStyle: aws.Bool(true),
			},
		}
	}

	return &s3_service.S3Service{
		Bucket:      envConfig.Bucket,
		CDNDomain:   CDNDomain,
		CDNProtocol: CDNProtocol,
		Config: &aws.Config{
			Credentials: credentials.NewStaticCredentials(envConfig.AccessKeyID, envConfig.AccessKey, ""),
			Endpoint:    aws.String(envConfig.Endpoint),
			Region:      aws.String(envConfig.Region),
		},
	}
}
