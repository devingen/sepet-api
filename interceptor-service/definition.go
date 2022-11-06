package is

import (
	"context"
	core "github.com/devingen/api-core"
	"github.com/devingen/sepet-api/dto"
)

// ISepetInterceptorService defines the functionality of the interceptors
type ISepetInterceptorService interface {
	Pre(ctx context.Context, req core.Request) (*dto.WebhookPreResponse, int, interface{})
	Final(ctx context.Context, req core.Request, responseBody interface{})
}
