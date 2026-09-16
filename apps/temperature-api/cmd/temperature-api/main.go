package main

import (
	"context"
	"net/http"
	"os/signal"
	"syscall"

	"temperature-api/internal/api/http/router"
	"temperature-api/internal/config"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	server := &http.Server{
		Addr:    cfg.Address,
		Handler: router.New(),
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()

	<-ctx.Done()

	server.Shutdown(context.Background())
}
