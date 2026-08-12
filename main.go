package main

import (
	"log"
	"net/http"
)

func main() {
	// Parse command-line flags
	cfg := parseFlags()

	// Prerequisite checks for running the application
	if !isGitInstalled() {
		log.Fatal("Git is not installed or not available in the system's PATH.")
	}

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
