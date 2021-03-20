package dto

// UploadFileResponse defines the response body structure of the upload bucket request
type UploadFileResponse struct {
	Locations []string `json:"locations,omitempty"`
}
