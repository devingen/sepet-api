package srvcont

import (
	"context"
	core "github.com/devingen/api-core"
	"github.com/devingen/api-core/log"
	"github.com/devingen/sepet-api/dto"
	"github.com/sirupsen/logrus"
	"net/http"
)

// GetFileList implements IServiceController interface
func (controller ServiceController) GetFileList(ctx context.Context, req core.Request) (*core.Response, error) {

	logger, err := log.Of(ctx)
	if err != nil {
		return nil, core.NewStatusError(http.StatusInternalServerError)
	}

	domain, path := GetDomainAndPath(req.Path, true)

	bucket, err := controller.DataService.FindBucketWithDomain(ctx, domain)
	if err != nil {
		return nil, err
	}
	if bucket == nil {
		return nil, core.NewError(http.StatusNotFound, "bucket-not-found")
	}

	bucketVersion := ""
	if *bucket.IsVersioningEnabled() {
		bucketVersion = core.StringValue(bucket.Version)
	}
	versionFromHeader, headerHasVersion := req.GetHeader("bucket-version")
	if headerHasVersion {
		if !core.BoolValue(bucket.IsVersioningEnabled()) {
			return nil, core.NewError(http.StatusBadRequest, "versioning-not-enabled")
		}
		bucketVersion = versionFromHeader
	}

	fileList, err := controller.FileService.GetFileList(ctx, bucket, bucketVersion, path, true)
	if err != nil {
		return nil, err
	}

	logger.WithFields(logrus.Fields{
		"sepet-domain": bucket.Domain,
	}).Debug("returning file list")

	return &core.Response{
		StatusCode: http.StatusOK,
		Body:       dto.GetFileListResponse{Results: fileList},
	}, nil
}
