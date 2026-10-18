package main

import (
	"context"
	"log"
	"os"

	"github.com/s1ntezc0der/event-horizon-backend/services/history/internal/app"
)

func main() {
	ctx := context.Background()
	application, err := app.New(ctx)
	if err != nil {
		log.Fatalf("failed to init history app: %v", err)
	}
	if err := application.RunUntilSignal(ctx); err != nil {
		log.Printf("history stopped: %v", err)
		os.Exit(1)
	}
}
