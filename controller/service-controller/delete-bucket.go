package srvcont

import (
	"context"
	core "github.com/devingen/api-core"
	"github.com/devingen/api-core/log"
	"github.com/sirupsen/logrus"
	"net/http"
)

// DeleteBucket implements IServiceController interface
func (controller ServiceController) DeleteBucket(ctx context.Context, req core.Request) (*core.Response, error) {

	_, interceptorStatusCode, interceptorError := controller.InterceptorService.Pre(ctx, req)
	if interceptorError != nil {
		return &core.Response{
			StatusCode: interceptorStatusCode,
			Body:       interceptorError,
		}, nil
	}

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

	err = controller.DataService.DeleteBucket(ctx, id)
	if err != nil {
		return nil, err
	}

	err = controller.FileService.DeleteEntireBucket(ctx, bucket)
	if err != nil {
		return nil, err
	}

	logger.WithFields(logrus.Fields{"bucket-id": id}).Debug("deleted bucket")

	controller.InterceptorService.Final(ctx, req, nil)

	return &core.Response{
		StatusCode: http.StatusNoContent,
	}, nil
}
