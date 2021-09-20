package search

import (
	"context"

	"bloglog.com/search/v1/internal/model"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/search"
)

type Adder func(ctx context.Context, page *model.Page) error

type Remover func(ctx context.Context, id string) error

type Searcher func(ctx context.Context, query string) ([]*model.Page, error)

type Client struct {
	client *search.Client
	index  *search.Index
}

func NewClient(appId, apiKey, indexName string) *Client {
	client := search.NewClient(appId, apiKey)
	index := client.InitIndex(indexName)

	return &Client{
		client: client,
		index:  index,
	}
}
