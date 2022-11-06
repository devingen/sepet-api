package srvcont

import (
	"context"
	core "github.com/devingen/api-core"
	"github.com/devingen/api-core/log"
	"github.com/sirupsen/logrus"
	"net/http"
)

// GetBucket implements IServiceController interface
func (controller ServiceController) GetBucket(ctx context.Context, req core.Request) (*core.Response, error) {

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

	logger.WithFields(logrus.Fields{
		"bucket-id":     bucket.ID,
		"bucket-domain": bucket.Domain,
	}).Debug("returning bucket")

	return &core.Response{
		StatusCode: http.StatusOK,
		Body:       bucket,
	}, nil
}
