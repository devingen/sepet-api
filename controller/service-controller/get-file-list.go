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
func (controller ServiceController) GetFileList(ctx context.Context, req core.Request) (interface{}, int, error) {

	logger, err := log.Of(ctx)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	domain, path := GetDomainAndPath(req.Path, true)

	bucket, err := controller.DataService.FindBucketWithDomain(ctx, domain)
	if err != nil {
		return nil, 0, err
	}
	if bucket == nil {
		return nil, 0, core.NewError(http.StatusNotFound, "bucket-not-found")
	}

	bucketVersion := bucket.Version
	versionFromHeader, headerHasVersion := req.GetHeader("Bucket-Version")
	if headerHasVersion {
		if !core.BoolValue(bucket.IsVersioningEnabled) {
			return nil, 0, core.NewError(http.StatusBadRequest, "versioning-not-enabled")
		}
		bucketVersion = core.String(versionFromHeader)
	}

	fileList, err := controller.FileService.GetFileList(ctx, bucket, *bucketVersion, path, true)
	if err != nil {
		return nil, 0, err
	}

	logger.WithFields(logrus.Fields{
		"sepet-domain": bucket.Domain,
	}).Debug("returning file list")

	return &dto.GetFileListResponse{Results: fileList}, http.StatusOK, nil
}
