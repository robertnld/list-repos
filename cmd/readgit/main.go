package main

import (
	"flag"
	"fmt"
	"list-repos/gitreader"
	"log"
	"os"
	"path/filepath"
)

// Structure for incoming parameters
type config struct {
	// The directory to scan for repositories
	ListDir string `default:"."`
}


func main() {
	cfg := config{}
	flag.StringVar(&cfg.ListDir, "list-dir", ".", "The directory to scan for repositories")
	flag.Parse()

	// Convert the provided directory path to an absolute path
	absPath, err := filepath.Abs(cfg.ListDir)
	if err != nil {
		println("Error getting absolute path:", err.Error())
		os.Exit(1)
	}
	cfg.ListDir = absPath
	log.Printf("Absolute repo directory: %s", cfg.ListDir)
	
	// Check if the specified directory is a Git repository
	if gitreader.IsGitRepository(cfg.ListDir) {
		println("This is a Git repository.")
	} else {
		println("This is NOT a Git repository.")
		os.Exit(1)
	}

	// Get commit object

	// Get the latest commit message for the Git repository at the specified path
	value, err := gitreader.GetLatestCommitMessage(cfg.ListDir)
	if err != nil {
		println("Error:", err.Error())
		return
	}
	fmt.Printf("Returned value: %s\n", value)
}
