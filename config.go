package main

import "flag"

// Structure for incoming parameters
type config struct {
	// Directory to list repositories from
	listDir string
	port    string
}

func parseFlags() config {
	listDir := flag.String(
		"dir",
		".",
		"Directory to list repositories from",
	)
	port := flag.String(
		"port",
		"8080",
		"Port to run the server on",
	)
	flag.Parse()

	return config{
		listDir: *listDir,
		port: *port,
	}
}
