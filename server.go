package main

import (
	"io/fs"
	"log"
	"net/http"
)


func newServer(cfg *config) (*http.ServeMux, error) {

	mux := http.NewServeMux()
	
	// Get the static files for serving
	staticFiles, err := fs.Sub(webFiles, "web/static")
	if err != nil {
		log.Fatal("Error accessing embedded static files:", err)
	}
	// Serve static files from the embedded filesystem
	mux.Handle(
		"/static/",
		http.StripPrefix(
			"/static/",
			http.FileServer(http.FS(staticFiles)),
		),
	)
	return mux, nil
}