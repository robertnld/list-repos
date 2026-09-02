package main

import (
	"list-repos/gitreader"
	"os"
	"os/exec"
	"strings"
)

type Repository struct {
	Name       string
	LastCommit string
	Branches   []Branch
}

type Branch struct {
	Name string
}

func listDirectories(path string) ([]string, error) {
	var directories []string

	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			directories = append(directories, entry.Name())
		}
	}

	return directories, nil
}

// List the Git repositories in the specified directory
func listGitRepositories(path string) ([]Repository, error) {
	directories, err := listDirectories(path)
	if err != nil {
		return nil, err
	}

	var gitRepositories []Repository
	for _, dir := range directories {
		fullPath := path + "/" + dir
		if gitreader.IsGitRepository(fullPath) {
			// Get last commit message
			LastCommit, err := gitreader.GetLatestCommitMessage(fullPath)
			if err != nil {
				return nil, err
			}

			// Get branches
			branches, err := listGitBranches(fullPath)
			if err != nil {
				return nil, err
			}
			var branchStructs []Branch
			for _, branch := range branches {
				branchStructs = append(branchStructs, Branch{Name: branch})
			}
			gitRepositories = append(gitRepositories, Repository{
				Name:       dir,
				LastCommit: strings.TrimSpace(string(LastCommit)),
				Branches:   branchStructs,
			})
		}
	}
	return gitRepositories, nil
}

// Show the branches of the repository in the specified directory
// TODO
func listGitBranches(repoPath string) ([]string, error) {

	// List the branches using the Git command
	cmd := exec.Command("git", "-C", repoPath, "branch", "--list")
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	branches := strings.Split(strings.TrimSpace(string(output)), "\n")
	for i, branch := range branches {
		branches[i] = strings.TrimSpace(branch)
	}
	return branches, nil
}
