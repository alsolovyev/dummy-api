package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/alsolovyev/dummy-api/internal/controller"
	"github.com/alsolovyev/dummy-api/internal/repository/filerepo"
	"github.com/alsolovyev/dummy-api/internal/usecase"
	"github.com/alsolovyev/dummy-api/pkg/httpserver"
	"github.com/alsolovyev/dummy-api/pkg/logger"
)

const (
	address = "127.0.0.1"
	port    = 8181
)

func main() {
	// Create a logger
	l := logger.New()
	l.Info("Launching the app")

	// Set up a channel to receive OS signals
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	// Create a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Initialize the repository layer
	filerepo := filerepo.New()

	// Initialize the usecase layer
	usecase := usecase.New(filerepo)

	// Initialize the controller layer
	r := controller.New(usecase.File)

	// Initialize the adapter layer
	// Create an HTTP server
	hs := httpserver.New(address, port, r)
	go func() {
		if err := hs.Run(); err != http.ErrServerClosed {
			l.Errorf("An error occurred while running HTTP server: %s", err.Error())
			stop <- syscall.SIGTERM
		}
	}()

	<-stop

	l.Info("Shutting down the app")

	// Shutdown the HTTP server
	if err := hs.Stop(ctx); err != nil {
		l.Errorf("Ann error occurred while shutting down HTTP server gracefully: %s", err.Error())
	}
}
