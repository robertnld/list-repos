/*
Show information for a specific Git repository.
The command does not use git.

Usage:

	readgit [OPTIONS] <path>

The options are	:

	    -latest string
			latest commit message

		-latest-commit
			all information of the latest commit

		-branch string
			show all branches

		-branch-remote string
			show all remote branches

		-branch-all
			show all branches (local and remote)

The app does not accept STDIN.
*/
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
	latestCommitMessage := flag.Bool(
		"latest",
		false,
		"print the latest commit message",
	)
	flag.Parse()

	listDir := "."
	if flag.NArg() > 0 {
		listDir = flag.Arg(0)
	}
	if flag.NArg() > 1 {
		println("Error: Too many arguments. Only one directory path is allowed.")
		os.Exit(1)
	}

	// Convert the provided directory path to an absolute path
	absPath, err := filepath.Abs(listDir)
	if err != nil {
		println("Error getting absolute path:", err.Error())
		os.Exit(1)
	}
	listDir = absPath
	log.Printf("Absolute repo directory: %s", listDir)

	// Check if the specified directory is a Git repository
	if gitreader.IsGitRepository(listDir) {
		println("This is a Git repository.")
	} else {
		println("This is NOT a Git repository.")
		os.Exit(1)
	}

	// Get the latest commit message for the Git repository at the specified path
	if *latestCommitMessage {

		value, err := gitreader.GetLatestCommitMessage(listDir)
		if err != nil {
			println("Error:", err.Error())
			return
		}
		fmt.Printf("Returned value: \n%s\n", value)
	}
}
