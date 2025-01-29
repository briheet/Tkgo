package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/briheet/tkgo/api"
	"github.com/briheet/tkgo/storage"
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
)

func main() {
	ctx := context.Background()

	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err)
		os.Exit(-1)
	}

	defer func() {
		if err := logger.Sync(); err != nil {
			log.Printf("Logger sync failed: %v", err)
			os.Exit(-1)
		}
	}()

	port := ":8080"
	storage := storage.NewStorage()

	c := InitCronScheduler(storage)
	defer c.Stop()

	server := api.NewServer(ctx, logger, port, storage)

	mux := http.NewServeMux()
	server.Serve(mux)

	logger.Info("Starting Server", zap.String("addr", port))

	if err := http.ListenAndServe(port, mux); err != nil {
		zap.Error(err)
	}
}

func InitCronScheduler(storage *storage.NonPresistentMap) *cron.Cron {
	c := cron.New()
	if _, err := c.AddFunc("@every 00h00m10s", func() { ResetTokens(storage) }); err != nil {
		log.Printf("Failed to schedule cron job: %v", err)
		os.Exit(-1)
	}
	c.Start()

	return c
}

func ResetTokens(storage *storage.NonPresistentMap) {
	storage.Mu.Lock()

	defer storage.Mu.Unlock()
	currentTime := time.Now()

	for userId, userData := range storage.Map {
		// Check if 24 hours have passed since last reset
		if currentTime.Sub(userData.TimeCreated) >= 24*time.Hour {
			fmt.Printf("Resetting tokens for user: %s\n", userId)

			// Reset all token counts to zero
			for token := range userData.TokenCount {
				userData.TokenCount[token] = 0
			}

			// Update the last reset timestamp
			userData.TimeCreated = currentTime
			storage.Map[userId] = userData // update the time
		}
	}
}
