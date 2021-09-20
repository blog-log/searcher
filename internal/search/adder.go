package search

import (
	"context"

	"bloglog.com/search/v1/internal/adapter"
	"bloglog.com/search/v1/internal/model"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/opt"
)

func (client *Client) Add(ctx context.Context, page *model.Page) error {
	// convert page to algoliaPage
	algoliaPage := adapter.PageToAlgolia(page)

	_, err := client.index.SaveObject(algoliaPage, opt.AutoGenerateObjectIDIfNotExist(false))
	if err != nil {
		return err
	}

	return nil
}
