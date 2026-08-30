package main

import (
	"flag"
	"fmt"
	"list-repos/gitreader"
	"log"
	"os"
	"path/filepath"
)


func main() {
	// Parameters
	listdir := flag.String("list-dir", ".", "The directory to scan for repositories")
	flag.Parse()

	// Convert the provided directory path to an absolute path
	absPath, err := filepath.Abs(*listdir)
	if err != nil {
		println("Error getting absolute path:", err.Error())
		os.Exit(1)
	}
	*listdir = absPath
	log.Printf("Absolute repo directory: %s", *listdir)
	
	// Check if the specified directory is a Git repository
	if gitreader.IsGitRepository(*listdir) {
		println("This is a Git repository.")
	} else {
		println("This is NOT a Git repository.")
		os.Exit(1)
	}

	// Get commit object

	// Get the latest commit message for the Git repository at the specified path
	value, err := gitreader.GetLatestCommitMessage(*listdir)
	if err != nil {
		println("Error:", err.Error())
		return
	}
	fmt.Printf("Returned value: %s\n", value)
}