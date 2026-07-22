package main

import (
	"embed"
	"flag"
	"fmt"
	"html/template"
	"net/http"
	"os"
)

// embed directive to include the web templates in the binary
//go:embed web/templates/* web/static/*
var webFiles embed.FS

var templates = template.Must(
	template.ParseFS(webFiles, "web/templates/*.html"),
)

func init() {
	// Define command-line flags
	flag.String("dir", ".", "Directory to list repositories from")

	isGitInstalled := isGitInstalled()
	if !isGitInstalled {
		fmt.Println("Git is not installed or not available in the system's PATH.")
		os.Exit(1)
	}
}

func main() {
	// Handle command-line flags
	flag.Parse()
	listDir := flag.Lookup("dir").Value.String()

	// Get the list of directories in the specified path
	directories, err := listDirectories(listDir)
	if err != nil {
		fmt.Println("Error listing directories:", err)
		return
	}

	// Filter out non-Git repositories
	var gitRepositories []string
	for _, dir := range directories {
		fullPath := listDir + "/" + dir
		if isGitRepository(fullPath) {
			gitRepositories = append(gitRepositories, dir)
		}
	}
	
	// Set up the HTTP server and routes
	mux := http.NewServeMux()
	// Serve static files from the embedded filesystem
	mux.Handle("/static/", http.FileServer(http.FS(webFiles)))

	data := struct {
		Title       string
		Directories []string
	}{
		Title:       "Repository List",
		Directories: gitRepositories,
	}

	// Handle the index route
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := templates.ExecuteTemplate(w, "index.html", data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Println("Starting server on :8080")
	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}