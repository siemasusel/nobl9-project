package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"stddevapi"
	"stddevapi/randomorg"
	"stddevapi/server"
	"syscall"

	"golang.org/x/exp/slog"
)

const addr = ":8080"

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	apiKey := os.Getenv("RANDOMORG_API_KEY")

	intBackend := randomorg.NewClient(apiKey)
	service := stddevapi.NewService(intBackend)
	httpHandler := server.NewHTTPHandler(service)

	runHTTPServer(ctx, httpHandler)
}

func runHTTPServer(ctx context.Context, handler http.Handler) {
	server := http.Server{Addr: addr, Handler: handler}
	slog.InfoContext(ctx, "Http server listening", "port", addr)

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.ErrorContext(ctx, "Http server failed")
		}
	}()

	<-ctx.Done()

	if err := server.Shutdown(context.Background()); err != nil {
		slog.ErrorContext(ctx, "Http server shutdown error")
	}

	slog.InfoContext(ctx, "Http server stopped")
}
