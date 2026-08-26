package main

import (
	"flag"
	"list-repos/gitreader"
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

	// Print the banner message as test
	line, err := gitreader.Banner()
	if err != nil {
		println("Error:", err.Error())
		return
	}
	println(line)

	// Check if the specified directory is a Git repository
	if gitreader.IsGitRepository(cfg.ListDir) {
		println("This is a Git repository.")
	} else {
		println("This is NOT a Git repository.")
	}

	// Get the latest commit message for the Git repository at the specified path
	commit_message, err := gitreader.GetLatestCommitMessage(cfg.ListDir)
	if err != nil {
		println("Error:", err.Error())
		return
	}
	println("Latest commit message:", commit_message)
}
