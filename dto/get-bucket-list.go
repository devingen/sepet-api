package dto

import "github.com/devingen/sepet-api/model"

type GetBucketListResponse struct {
	Results []*model.Bucket `json:"results"`
}
