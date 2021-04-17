package srvcont

import (
	"context"
	core "github.com/devingen/api-core"
	"github.com/devingen/api-core/log"
	"github.com/devingen/sepet-api/dto"
	"github.com/devingen/sepet-api/model"
	"net/http"
)

// GetBuckets implements IServiceController interface
func (controller ServiceController) GetBuckets(ctx context.Context, req core.Request) (interface{}, int, error) {
	logger, err := log.Of(ctx)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	buckets, err := controller.DataService.FindBuckets(ctx, model.BucketStatusActive)

	logger.Debug("returning bucket list")
	return &dto.GetBucketListResponse{Results: buckets}, http.StatusOK, err
}
