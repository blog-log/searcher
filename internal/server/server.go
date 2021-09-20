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

	mux := mux.NewRouter().StrictSlash(true)

	// setup SearchHandler
	handler := handler.NewSearchHandler(cfg)

	mux.HandleFunc("/", handler.Create).Methods("POST")
	mux.HandleFunc("/{id}", handler.Remove).Methods("DELETE")
	mux.HandleFunc("/search", handler.Search).Methods("POST")

	fmt.Println("Server started at port 8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		return err
	}

	return nil
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to my web server!</h1>"))
}
