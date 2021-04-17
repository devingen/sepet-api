package srvcont

import (
	"context"
	core "github.com/devingen/api-core"
	"github.com/devingen/api-core/log"
	"github.com/sirupsen/logrus"
	"net/http"
)

// DeleteFile implements IServiceController interface
func (controller ServiceController) DeleteFile(ctx context.Context, req core.Request) (interface{}, int, error) {

	logger, err := log.Of(ctx)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	domain, path := GetDomainAndPath(req.Path, false)

	//domain := req.PathParameters["pathToDelete"]
	//if domain == "" {
	//	return "", http.StatusBadRequest, core.NewError(http.StatusBadRequest, "path-missing")
	//}
	// should not trim the / in case it's a folder
	//path := GetFilePath(req.Path, domain, false)

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

	err = controller.FileService.DeleteFile(ctx, bucket, *bucketVersion, path)
	if err != nil {
		return nil, 0, err
	}

	logger.WithFields(logrus.Fields{
		"sepet-domain": bucket.Domain,
	}).Debug("deleted file")

	return nil, http.StatusNoContent, nil
}
