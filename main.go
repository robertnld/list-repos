package main

import (
	"log"
	"net/http"
)

func main() {
	// Parse command-line flags
	cfg := parseFlags()

	// Set up the HTTP server
	mux, err := newServer(cfg)
	if err != nil {
		log.Fatalf("create server: %v", err)
	}

	// Start the HTTP server
	log.Println("Starting server on port", cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, mux); err != nil {
		log.Fatalf("start server: %v", err)
	}
}
