package main

import (
	"os"
	"os/exec"
	"strings"
)

type Repository struct {
	Name     string
	Branches []Branch
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

// Test that Git is installed and available in the system's PATH
func isGitInstalled() bool {
	_, err := os.Stat("/usr/bin/git")
	return err == nil
}

// Test that directory is a Git repository by checking for the presence of a .git folder
func isGitRepository(path string) bool {
	gitDir := path + "/.git"
	info, err := os.Stat(gitDir)
	if err != nil {
		return false
	}
	return info.IsDir()
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
		if isGitRepository(fullPath) {
			gitRepositories = append(gitRepositories, Repository{Name: dir})
		}
	}
	return gitRepositories, nil
}

// Show the branches of the repository in the specified directory
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