package srvcont

import (
	"context"
	core "github.com/devingen/api-core"
	core_dto "github.com/devingen/api-core/dto"
	"github.com/devingen/api-core/log"
	"github.com/devingen/sepet-api/dto"
	"github.com/sirupsen/logrus"
	"net/http"
)

// UpdateBucket implements IServiceController interface
func (controller ServiceController) UpdateBucket(ctx context.Context, req core.Request) (interface{}, int, error) {

	logger, err := log.Of(ctx)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	id, hasID := req.PathParameters["id"]
	if !hasID {
		return nil, 0, core.NewError(http.StatusBadRequest, "id-missing")
	}

	var body dto.UpdateBucketRequest
	err = req.AssertBody(&body)
	if err != nil {
		return nil, 0, err
	}

	bucket, err := controller.DataService.FindBucketWithID(ctx, id)
	if err != nil {
		return nil, 0, err
	}

	if bucket == nil {
		return nil, 0, core.NewError(http.StatusNotFound, "bucket-not-found")
	}

	if body.Domain != nil && *body.Domain != "" {
		// check if another bucket has the same domain
		existingBucketWithSameDomain, err := controller.DataService.FindBucketWithDomain(ctx, *body.Domain)
		if err != nil {
			return nil, 0, err
		}
		if existingBucketWithSameDomain != nil && id != existingBucketWithSameDomain.ID.Hex() {
			return nil, 0, core.NewError(http.StatusConflict, "domain-is-already-being-used")
		}
		bucket.Domain = body.Domain
	}

	bucket.IndexPagePath = body.IndexPagePath
	bucket.ErrorPagePath = body.ErrorPagePath
	bucket.IsCacheEnabled = body.IsCacheEnabled
	bucket.IsVersioningEnabled = body.IsVersioningEnabled
	bucket.Status = body.Status
	bucket.Version = body.Version
	bucket.VersionIdentifier = body.VersionIdentifier
	bucket.CORSConfigs = body.CORSConfigs

	updatedAt, revision, err := controller.DataService.UpdateBucket(ctx, bucket)
	if err != nil {
		return nil, 0, err
	}

	logger.WithFields(logrus.Fields{
		"bucket-id":     bucket.ID,
		"bucket-domain": bucket.Domain,
	}).Debug("updated bucket")

	return core_dto.UpdateEntryResponse{
		ID:        id,
		UpdatedAt: *updatedAt,
		Revision:  revision,
	}, http.StatusOK, nil
}
