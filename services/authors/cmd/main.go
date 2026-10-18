package main

import (
	"context"
	"log"
	"os"

	"github.com/s1ntezc0der/event-horizon-backend/services/authors/internal/app"
)

func main() {
	ctx := context.Background()
	application, err := app.New(ctx)
	if err != nil {
		log.Fatalf("failed to init authors app: %v", err)
	}
	if err := application.RunUntilSignal(ctx); err != nil {
		log.Printf("authors stopped: %v", err)
		os.Exit(1)
	}
}
