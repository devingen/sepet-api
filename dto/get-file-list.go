package dto

// GetFileListResponse defines the response body structure of the create bucket request
type GetFileListResponse struct {
	Files []string `json:"files"`
}
