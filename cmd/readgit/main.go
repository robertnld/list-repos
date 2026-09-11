/*
Show information for a specific Git repository.
The command does not use git.

Usage:

	readgit [flags] <path>

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
	"os"
)

var repoDir = flag.String(
	"path",
	".",
	"directory path of the Git repository",
)
var latestCommitMessage = flag.Bool(
	"latest",
	false,
	"print the latest commit message",
)
var allBranches = flag.Bool(
	"branch-all",
	false,
	"print all branches (local and remote)",
)

func main() {
	flag.Parse()


	// Only if the directory has a Git repository
	isGitRepo, err := gitreader.IsGitRepository(*repoDir)
	if err != nil {
		println("Error:", err.Error())
		os.Exit(1)
	}
	if !isGitRepo {
		println("This is NOT a Git repository.")
		os.Exit(1)
	}	


	// Get the latest commit message for the Git repository at the specified path
	if *latestCommitMessage {

		value, err := gitreader.GetLatestCommitMessage(*repoDir)
		if err != nil {
			println("Error:", err.Error())
			return
		}
		fmt.Printf("Returned value: \n%s\n", value)
	}


	// Print all branches if the flag is set
	if *allBranches {
		branches, err := gitreader.GetAllBranches(*repoDir)
		if err != nil {
			println("Error:", err.Error())
			return
		}

		for _, branch := range branches {
			fmt.Println(branch)
		}
	}
}
