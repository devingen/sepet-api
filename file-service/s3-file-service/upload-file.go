package s3fs

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	core "github.com/devingen/api-core"
	"github.com/devingen/api-core/log"
	"github.com/devingen/sepet-api/model"
	"io"
	"mime/multipart"
	"net/http"
	"strings"
)

// UploadFile implements ISepetService interface
func (s3Service S3Service) UploadFile(ctx context.Context, bucket *model.Bucket, bucketVersion string, files map[string]multipart.File) ([]string, error) {

	logger, err := log.Of(ctx)
	if err != nil {
		return nil, err
	}

	sess := session.New(s3Service.Config)
	uploader := s3manager.NewUploader(sess)

	locations := make([]string, 0)
	for fileKey, file := range files {
		var body io.Reader = file
		if file == nil {
			// the inputs without files are empty folders to be created
			isFolder := fileKey[len(fileKey)-1:] == "/"
			if !isFolder {
				return nil, core.NewError(http.StatusBadRequest, "folder-names-must-end-with-/")
			}
			// create empty content for create folder request
			body = strings.NewReader(" ")
		} else {
			defer func() {
				err := file.Close()
				if err != nil {
					logger.Error(err)
				}
			}()
		}

		_, err := uploader.Upload(&s3manager.UploadInput{
			Bucket: aws.String(s3Service.Bucket),
			Key:    aws.String(GetFilePath(core.StringValue(bucket.Folder), bucketVersion, fileKey)),
			Body:   body,
		})
		if err != nil && err.(awserr.RequestFailure).StatusCode() != 409 {
			return nil, err
		}
		locations = append(locations, GenerateSepetCDNURL(s3Service.CDNDomain, s3Service.CDNProtocol, core.StringValue(bucket.Domain), fileKey))
	}
	return locations, nil
}

// GenerateSepetCDNURL generates the Bucket CDN URL of the uploaded file with the Bucket name.
// e.g. https://70l9ad53s5.sepet.devingen.io/images/selimiye.jpg
func GenerateSepetCDNURL(CDNDomain, CDNProtocol string, sepetName, key string) string {
	return CDNProtocol + "://" + sepetName + "." + CDNDomain + "/" + key
}

// GetFilePath generates the exact file path for the given Bucket's folder
// e.g. folder-for-70l9ad53s5/0.0.2/images/selimiye.jpg
func GetFilePath(folder, version, path string) string {
	filePath := folder + "/" + version + "/" + path
	return filePath
}
