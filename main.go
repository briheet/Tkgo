package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/briheet/tkgo/api"
	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err)
		os.Exit(-1)
	}
	defer logger.Sync()

	port := ":8080"
	ctx := context.Background()
	server := api.NewServer(ctx, logger, port)

	mux := http.NewServeMux()
	server.Serve(mux)

	logger.Info("Starting Server", zap.String("addr", port))

	if err := http.ListenAndServe(port, mux); err != nil {
		zap.Error(err)
	}
}
