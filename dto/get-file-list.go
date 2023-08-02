package dto

import "github.com/devingen/sepet-api/model"

// GetFileListResponse defines the response body structure of the create bucket request
type GetFileListResponse struct {
	Results []model.File `json:"results"`
}
