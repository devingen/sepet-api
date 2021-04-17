package dto

// UpdateBucketRequest defines the request body structure of the update bucket request
type UpdateBucketRequest struct {
	Domain              *string `json:"domain,omitempty" validate:"omitempty,min=5,max=32,alphanum"`
	Region              *string `json:"region,omitempty" validate:"omitempty,oneof='eu-central-1'"`
	IndexPagePath       *string `json:"indexPagePath,omitempty"`
	ErrorPagePath       *string `json:"errorPagePath,omitempty"`
	IsCacheEnabled      *bool   `json:"isCacheEnabled,omitempty"`
	IsVersioningEnabled *bool   `json:"isVersioningEnabled,omitempty"`
	Status              *string `json:"status,omitempty" validate:"omitempty,oneof='active'"`
	Version             *string `json:"version,omitempty" validate:"omitempty,min=1"`
	VersionIdentifier   *string `json:"versionIdentifier,omitempty" validate:"omitempty,oneof='header' 'path'"`
}
