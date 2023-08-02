package srvcont

import (
	"context"
	core "github.com/devingen/api-core"
	"github.com/devingen/api-core/log"
	"github.com/devingen/sepet-api/dto"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

// GetBucketVersionList implements IServiceController interface
func (controller ServiceController) GetBucketVersionList(ctx context.Context, req core.Request) (*core.Response, error) {

	logger, err := log.Of(ctx)
	if err != nil {
		return nil, core.NewStatusError(http.StatusInternalServerError)
	}

	id, hasID := req.PathParameters["id"]
	if !hasID {
		return nil, core.NewError(http.StatusBadRequest, "id-missing")
	}

	bucket, err := controller.DataService.FindBucketWithID(ctx, id)
	if err != nil {
		return nil, err
	}

	if bucket == nil {
		return nil, core.NewError(http.StatusNotFound, "bucket-not-found")
	}

	fileList, err := controller.FileService.GetFileList(ctx, bucket, "", "", true)
	if err != nil {
		return nil, err
	}

	logger.WithFields(logrus.Fields{
		"sepet-domain": bucket.Domain,
	}).Debug("returning file list")

	for i := range fileList {
		fileList[i].Path = strings.TrimSuffix(fileList[i].Path, "/")
	}

	return &core.Response{
		StatusCode: http.StatusOK,
		Body:       dto.GetFileListResponse{Results: fileList},
	}, nil
}
