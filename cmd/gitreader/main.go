package main

import (
	"flag"
	"list-repos/gitreader"
)


func main() {
	// Parameters
	listdir := flag.String("list-dir", ".", "The directory to scan for repositories")
	flag.Parse()

	// Check if the specified directory is a Git repository
	if gitreader.IsGitRepository(*listdir) {
		println("This is a Git repository.")
	} else {
		println("This is NOT a Git repository.")
	}

	// Get the latest commit message for the Git repository at the specified path
	commit_message, err := gitreader.GetLatestCommitMessage(*listdir)
	if err != nil {
		println("Error:", err.Error())
		return
	}
	println("Latest commit message:", commit_message)
}