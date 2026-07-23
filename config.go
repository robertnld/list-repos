package main

import "flag"

type config struct {
	// Directory to list repositories from
	listDir string
}

func parseFlags() config {
	listDir := flag.String(
		"dir",
		".",
		"Directory to list repositories from",
	)
	flag.Parse()

	return config{
		listDir: *listDir,
	}
}