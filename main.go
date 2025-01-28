package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/briheet/tkgo/api"
	"github.com/briheet/tkgo/storage"
	"go.uber.org/zap"
)

func main() {
	ctx := context.Background()

	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err)
		os.Exit(-1)
	}
	defer logger.Sync()

	port := ":8080"
	storage := storage.NewStorage()

	server := api.NewServer(ctx, logger, port, storage)

	mux := http.NewServeMux()
	server.Serve(mux)

	logger.Info("Starting Server", zap.String("addr", port))

	if err := http.ListenAndServe(port, mux); err != nil {
		zap.Error(err)
	}
}
