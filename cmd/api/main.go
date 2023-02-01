package main

import "github.com/alsolovyev/dummy-api/pkg/logger"

func main() {
	// Create a logger
	l := logger.New()
	l.Info("Launching the app")
}
