package main

import (
	"fmt"
	"net/http"

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

	// Initialize the adapter layer
	// Create an HTTP server
	m := http.NewServeMux()
	m.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello\n")
	}))
	hs := httpserver.New(address, port, m)
	if err := hs.Run(); err != http.ErrServerClosed {
		l.Errorf("An error occurred while running HTTP server: %s", err.Error())
	}
}
