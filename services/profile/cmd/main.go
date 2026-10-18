package main

import (
	"context"
	"log"
	"os"

	"github.com/s1ntezc0der/event-horizon-backend/services/profile/internal/app"
)

func main() {
	ctx := context.Background()
	application, err := app.New(ctx)
	if err != nil {
		log.Fatalf("failed to init profile app: %v", err)
	}

	if err := application.RunUntilSignal(ctx); err != nil {
		log.Printf("profile stopped with error: %v", err)
		os.Exit(1)
	}
}
