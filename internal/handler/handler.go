package handler

import (
	"bloglog.com/search/v1/internal/config"
	"bloglog.com/search/v1/internal/search"
)

type SearchHandler struct {
	adder    search.Adder
	remover  search.Remover
	searcher search.Searcher
}

func NewSearchHandler(cfg *config.Algolia) *SearchHandler {
	client := search.NewClient(cfg.AppId, cfg.ApiKey, cfg.IndexName)

	return &SearchHandler{
		adder:    client.Add,
		remover:  client.Remove,
		searcher: client.Search,
	}
}
