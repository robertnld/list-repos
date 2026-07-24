package main

import (
	"embed"
	"fmt"
	"html/template"
	"io/fs"
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

func newServer(cfg config) (*http.ServeMux, error) {

	mux := http.NewServeMux()

	// Get the static files for serving
	staticFiles, err := fs.Sub(webFiles, "web/static")
	if err != nil {
		return nil, fmt.Errorf(
			"access embedded static files: %w", err)
	}

	// Serve static files from the embedded filesystem
	mux.Handle(
		"/static/",
		http.StripPrefix(
			"/static/",
			http.FileServer(http.FS(staticFiles)),
		),
	)

	// Handle the index route
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		// Get the list of repositories in the specified path
		gitRepositories, err := listGitRepositories(cfg.listDir)
		if err != nil {
			log.Printf("Error listing Git repositories: %v", err)
			http.Error(
				w,
				"Error listing Git repositories",
				http.StatusInternalServerError,
			)
			return
		}

		data := pageData{
			Title:        pageTitle,
			Repositories: gitRepositories,
		}

		if err := templates.ExecuteTemplate(w, "index.html", data); err != nil {
			log.Printf("render index template: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	})

	return mux, nil
}
