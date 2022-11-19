package dto

// CreateBucketRequest defines the request body structure of the create bucket request
type CreateBucketRequest struct {
	Domain            *string            `json:"domain,omitempty" validate:"bucket-domain"`
	Region            *string            `json:"region,omitempty" validate:"oneof='eu-central-1'"`
	IndexPagePath     *string            `json:"indexPagePath,omitempty"`
	ErrorPagePath     *string            `json:"errorPagePath,omitempty"`
	IsCacheEnabled    *bool              `json:"isCacheEnabled,omitempty"`
	VersionIdentifier *string            `json:"versionIdentifier,omitempty"`
	ResponseHeaders   *map[string]string `json:"responseHeaders,omitempty" bson:"responseHeaders,omitempty"`
}
