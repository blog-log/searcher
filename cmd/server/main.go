package main

import (
	"context"
	"log"

	"bloglog.com/search/v1/internal/config"
	serverv1 "bloglog.com/search/v1/internal/server"
)

func main() {
	ctx := context.Background()

	cfg, err := config.Load()
	if err != nil {
		log.Fatal("error loading config", err)
	}

	server := serverv1.NewSearchServer()

	server.Run(ctx, cfg)
}
