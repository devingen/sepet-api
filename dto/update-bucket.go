package dto

import "github.com/devingen/sepet-api/model"

// UpdateBucketRequest defines the request body structure of the update bucket request
type UpdateBucketRequest struct {
	Domain              *string             `json:"domain,omitempty" validate:"omitempty,bucket-domain"`
	Region              *string             `json:"region,omitempty" validate:"omitempty,oneof='eu-central-1'"`
	IndexPagePath       *string             `json:"indexPagePath,omitempty"`
	ErrorPagePath       *string             `json:"errorPagePath,omitempty"`
	IsCacheEnabled      *bool               `json:"isCacheEnabled,omitempty"`
	IsVersioningEnabled *bool               `json:"isVersioningEnabled,omitempty"`
	Status              *string             `json:"status,omitempty" validate:"omitempty,oneof='active'"`
	Version             *string             `json:"version,omitempty" validate:"omitempty,min=1"`
	VersionIdentifier   *string             `json:"versionIdentifier,omitempty" validate:"omitempty,oneof='header' 'path'"`
	CORSConfigs         *[]model.CORSConfig `json:"corsConfigs,omitempty" bson:"corsConfig,omitempty"`
	ResponseHeaders     *map[string]string  `json:"responseHeaders,omitempty" bson:"responseHeaders,omitempty"`
}
