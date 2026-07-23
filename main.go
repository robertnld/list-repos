package main

import (
	"embed"
	"fmt"
	"html/template"
	"io/fs"
	"net/http"
	"os"
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

	if !isGitInstalled() {
		fmt.Println("Git is not installed or not available in the system's PATH.")
		os.Exit(1)
	}
	
	// Get the static files for serving
	staticFiles, err := fs.Sub(webFiles, "web/static")
	if err != nil {
		fmt.Println("Error accessing embedded static files:", err)
		return
	}

	// Set up the HTTP server and routes
	mux := http.NewServeMux()
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
		// Get the list of directories in the specified path
		directories, err := listDirectories(cfg.listDir)
		if err != nil {
			http.Error(
				w,
				"Error listing directories: "+err.Error(),
				http.StatusInternalServerError,
			)
			fmt.Println("Error listing directories:", err)
			return
		}

		// Filter out non-Git repositories
		var gitRepositories []string
		for _, dir := range directories {
			fullPath := cfg.listDir + "/" + dir
			if isGitRepository(fullPath) {
				gitRepositories = append(gitRepositories, dir)
			}
		}

		data := pageData{
			Title:        pageTitle,
			Repositories: gitRepositories,
		}

		if err := templates.ExecuteTemplate(w, "index.html", data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Println("Starting server on :8080")
	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
