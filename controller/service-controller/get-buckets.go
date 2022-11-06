package srvcont

import (
	"context"
	core "github.com/devingen/api-core"
	"github.com/devingen/api-core/log"
	"github.com/devingen/sepet-api/dto"
	"net/http"
)

// GetBuckets implements IServiceController interface
func (controller ServiceController) GetBuckets(ctx context.Context, req core.Request) (*core.Response, error) {
	logger, err := log.Of(ctx)
	if err != nil {
		return nil, core.NewStatusError(http.StatusInternalServerError)
	}

	query, _, iStatusCode, iErr := PreGetQueryEnhance(controller, ctx, req)
	if iErr != nil {
		return &core.Response{
			StatusCode: iStatusCode,
			Body:       iErr,
		}, nil
	}

	buckets, err := controller.DataService.FindBuckets(ctx, query)

	logger.Debug("returning bucket list")
	return &core.Response{
		StatusCode: http.StatusOK,
		Body:       dto.GetBucketListResponse{Results: buckets},
	}, nil
}
