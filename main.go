package main

import (
	"context"
	"errors"
	"go_server/handler"
	"go_server/helper"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	log.Println("Local server started in on port https://localhost:3072")
	helper.GetCurrentDateTime()
	// Create a new HTTP server instance
	srv := &http.Server{Addr: ":3072", Handler: nil} // Handler is nil here; use http.HandleFunc

	// Use http.HandleFunc to register the handler:
	http.HandleFunc("/graphql", handler.Handler)

	// Create a context that can be cancelled
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // Cancel the context when the program exits

	// Start the server in a goroutine so it doesn't block
	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(http.ErrServerClosed, err) { // Use the srv instance
			log.Fatalf("ListenAndServe error: %v", err)
		}
	}()

	// Handle signals for graceful shutdown
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	// Wait for a signal
	<-sigchan

	log.Println("Shutting down server...")

	// Gracefully shut down the server using the srv instance and the context
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown error: %v", err)
	}

	log.Println("Server gracefully shut down.")
}
