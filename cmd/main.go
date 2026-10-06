package main

import (
	"context"
	"errors"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/ppechkurov/sobesrazborgoogle/internal/channels"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	err := channels.Run(ctx)
	if err != nil && !errors.Is(err, context.Canceled) {
		log.Printf("run: %v", err)
		os.Exit(1)
	}
}
