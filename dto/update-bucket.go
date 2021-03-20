package dto

import "github.com/devingen/sepet-api/model"

// UpdateBucketRequest defines the request body structure of the update bucket request
type UpdateBucketRequest struct {
	Domain              string             `json:"domain,omitempty"`
	IndexPagePath       string             `json:"indexPagePath,omitempty"`
	ErrorPagePath       string             `json:"errorPagePath,omitempty"`
	IsCacheEnabled      bool               `json:"isCacheEnabled,omitempty"`
	IsVersioningEnabled bool               `json:"isVersioningEnabled,omitempty"`
	Status              model.BucketStatus `json:"status,omitempty" bson:"status,omitempty"`
}
