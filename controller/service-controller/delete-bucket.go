package srvcont

import (
	"context"
	core "github.com/devingen/api-core"
	"github.com/devingen/api-core/log"
	"github.com/sirupsen/logrus"
	"net/http"
)

// DeleteBucket implements IServiceController interface
func (controller ServiceController) DeleteBucket(ctx context.Context, req core.Request) (interface{}, int, error) {

	logger, err := log.Of(ctx)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	id, hasID := req.PathParameters["id"]
	if !hasID {
		return nil, 0, core.NewError(http.StatusBadRequest, "id-missing")
	}

	bucket, err := controller.DataService.FindBucketWithID(ctx, id)
	if err != nil {
		return nil, 0, err
	}

	if bucket == nil {
		return nil, 0, core.NewError(http.StatusNotFound, "bucket-not-found")
	}

	err = controller.DataService.DeleteBucket(ctx, id)
	if err != nil {
		return nil, 0, err
	}

	err = controller.FileService.DeleteEntireBucket(ctx, bucket)
	if err != nil {
		return nil, 0, err
	}

	logger.WithFields(logrus.Fields{"bucket-id": id}).Debug("deleted bucket")

	return nil, http.StatusNoContent, nil
}
