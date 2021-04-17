package dto

// CreateBucketRequest defines the request body structure of the create bucket request
type CreateBucketRequest struct {
	Domain              *string `json:"domain,omitempty" validate:"min=5,max=32,alphanum"`
	Region              *string `json:"region,omitempty" validate:"oneof='eu-central-1'"`
	IndexPagePath       *string `json:"indexPagePath,omitempty"`
	ErrorPagePath       *string `json:"errorPagePath,omitempty"`
	IsCacheEnabled      *bool   `json:"isCacheEnabled,omitempty"`
	IsVersioningEnabled *bool   `json:"isVersioningEnabled,omitempty"`
}
