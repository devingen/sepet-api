package s3fs

import (
	"github.com/aws/aws-sdk-go/aws"
)

// S3Service implements ISepetService interface with database connection
type S3Service struct {
	Bucket      string
	CDNDomain   string
	CDNProtocol string
	Config      *aws.Config
}
