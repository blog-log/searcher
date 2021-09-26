package server

import (
	"context"
	"fmt"
	"net/http"

	"bloglog.com/search/v1/internal/config"
	"bloglog.com/search/v1/internal/handler"
	"github.com/gorilla/mux"
)

type Server interface {
	Run(ctx context.Context) error
}

type SearchServer struct{}

func NewSearchServer() *SearchServer {
	return &SearchServer{}
}

func (s *SearchServer) Run(ctx context.Context, cfg *config.Algolia) error {

	rootRouter := mux.NewRouter().StrictSlash(true)

	documentRouter := rootRouter.PathPrefix("/document").Subrouter()

	// setup SearchHandler
	handler := handler.NewSearchHandler(cfg)

	// define routes
	documentRouter.HandleFunc("/", handler.Create).Methods("POST")
	documentRouter.HandleFunc("/{id}", handler.Remove).Methods("DELETE")
	documentRouter.HandleFunc("/search", handler.Search).Methods("POST")

	fmt.Println("Server started at port 8080")
	if err := http.ListenAndServe(":8080", rootRouter); err != nil {
		return err
	}

	return nil
}
