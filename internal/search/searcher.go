package search

import (
	"context"

	"bloglog.com/search/v1/internal/adapter"
	"bloglog.com/search/v1/internal/model"
)

func (client *Client) Search(ctx context.Context, query string) ([]*model.Page, error) {
	res, err := client.index.Search(query)
	if err != nil {
		return nil, err
	}

	var algoliaPages []*model.AlgoliaPage
	if err := res.UnmarshalHits(&algoliaPages); err != nil {
		return nil, err
	}

	// convert all algoliaPages to pages
	var pages []*model.Page
	for _, algoliaPage := range algoliaPages {
		pages = append(pages, adapter.AlgoliaToPage(algoliaPage))
	}

	return pages, nil
}
