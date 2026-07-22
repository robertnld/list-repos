package main

import (
	"embed"
	"flag"
	"fmt"
	"html/template"
	"net/http"
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

	mux := http.NewServeMux()
	// Serve static files from the embedded filesystem
	mux.Handle("/static/", http.FileServer(http.FS(webFiles)))

	// Handle the index route
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := templates.ExecuteTemplate(w, "index.html", directories)
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

