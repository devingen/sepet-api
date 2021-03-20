package srvcont

import (
	"context"
	core "github.com/devingen/api-core"
	"github.com/devingen/api-core/log"
	"github.com/devingen/api-core/util"
	"github.com/devingen/sepet-api/dto"
	"github.com/devingen/sepet-api/model"
	"github.com/sirupsen/logrus"
	"net/http"
)

// CreateBucket implements IServiceController interface
func (controller ServiceController) CreateBucket(ctx context.Context, req core.Request) (interface{}, int, error) {

	logger, err := log.Of(ctx)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	var body dto.CreateBucketRequest
	err = req.AssertBody(&body)
	if err != nil {
		return nil, 0, err
	}

	// check if another bucket has the same domain
	existingBucketWithSameDomain, err := controller.DataService.FindBucketWithDomain(ctx, body.Domain)
	if err != nil {
		return nil, 0, err
	}
	if existingBucketWithSameDomain != nil {
		return nil, 0, core.NewError(http.StatusConflict, "domain-is-already-being-used")
	}

	// generate a folder name until found unique
	folder := util.GenerateRandomString(16)
	for {
		// check if another bucket has the same folder
		existingBucketWithSameFolder, err := controller.DataService.FindBucketWithFolder(ctx, folder)
		if err != nil {
			return nil, 0, err
		}
		if existingBucketWithSameFolder == nil {
			// break if there is no bucket with the folder name
			break
		}

		// generate a new folder name
		folder = util.GenerateRandomString(16)
	}

	bucket := &model.Bucket{
		Domain:              body.Domain,
		Folder:              folder,
		Version:             "default",
		IndexPagePath:       body.IndexPagePath,
		ErrorPagePath:       body.ErrorPagePath,
		IsCacheEnabled:      body.IsCacheEnabled,
		IsVersioningEnabled: body.IsVersioningEnabled,
		Status:              model.BucketStatusActive,
	}

	bucket, err = controller.DataService.CreateBucket(ctx, bucket)
	if err != nil {
		return nil, 0, err
	}

	logger.WithFields(logrus.Fields{
		"bucket-id":     bucket.ID,
		"bucket-domain": bucket.Domain,
	}).Debug("created bucket")

	return &bucket, http.StatusOK, err
}
