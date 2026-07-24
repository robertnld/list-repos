package main

import (
	"embed"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// embed directive to include the web templates in the binary
//
//go:embed web/templates/* web/static/*
var webFiles embed.FS

var templates = template.Must(
	template.ParseFS(webFiles, "web/templates/*.html"),
)
const pageTitle = "Repository List"
type pageData struct {
    Title        string
    Repositories []string
}

func main() {
	// Parse command-line flags
	cfg := parseFlags()

	// Prerequisite checks for running the application
	if !isGitInstalled() {
		log.Fatal("Git is not installed or not available in the system's PATH.")
	}

	// Set up the HTTP server
	mux, err := newServer(&cfg)
	if err != nil {
		log.Fatal("Error creating server:", err)
	}

	
	// Handle the index route
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		
		// Get the list of repositories in the specified path
		gitRepositories, err := listGitRepositories(cfg.listDir)
		if err != nil {
			http.Error(
				w,
				"Error listing Git repositories: "+err.Error(),
				http.StatusInternalServerError,
			)
			fmt.Println("Error listing Git repositories:", err)
			return
		}

		data := pageData{
			Title:        pageTitle,
			Repositories: gitRepositories,
		}

		if err := templates.ExecuteTemplate(w, "index.html", data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
