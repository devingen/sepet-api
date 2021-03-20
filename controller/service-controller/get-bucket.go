package srvcont

import (
	"context"
	core "github.com/devingen/api-core"
	"github.com/devingen/api-core/log"
	"github.com/sirupsen/logrus"
	"net/http"
)

// GetBucket implements IServiceController interface
func (controller ServiceController) GetBucket(ctx context.Context, req core.Request) (interface{}, int, error) {

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

	logger.WithFields(logrus.Fields{
		"bucket-id":     bucket.ID,
		"bucket-domain": bucket.Domain,
	}).Debug("returning bucket")

	return bucket, http.StatusOK, nil
}
