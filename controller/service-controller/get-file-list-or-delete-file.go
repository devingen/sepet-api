package srvcont

import (
	"context"
	core "github.com/devingen/api-core"
	"github.com/devingen/api-core/log"
	"github.com/sirupsen/logrus"
	"net/http"
)

// GetFileListOrDeleteFile implements IServiceController interface
func (controller ServiceController) GetFileListOrDeleteFile(ctx context.Context, req core.Request) (*core.Response, error) {

	if req.HTTPMethod == http.MethodGet {
		return controller.GetFileList(ctx, req)
	} else if req.HTTPMethod == http.MethodDelete {
		return controller.DeleteFile(ctx, req)
	}

	logger, err := log.Of(ctx)
	if err != nil {
		return nil, core.NewStatusError(http.StatusInternalServerError)
	}

	logger.WithFields(logrus.Fields{
		"path":   req.Path,
		"method": req.HTTPMethod,
	}).Debug("couldn't find a controller for the method")

	return nil, core.NewStatusError(http.StatusInternalServerError)
}
