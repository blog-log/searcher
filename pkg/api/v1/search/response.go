package search

import "bloglog.com/search/v1/internal/model"

type Response struct {
	Documents []*model.Page `json:"documents"`
}
