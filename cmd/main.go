package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/ppechkurov/sobesrazborgoogle/internal/channels"

	_ "net/http/pprof"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)

	go func() {
		http.ListenAndServe("127.0.0.1:6060", nil)
	}()

	err := channels.Run(ctx)
	if err != nil && !errors.Is(err, context.Canceled) {
		log.Printf("run: %v", err)
		stop()
		os.Exit(1)
	}
}
