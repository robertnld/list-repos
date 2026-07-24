package main

import "os"

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
func listGitRepositories(path string) ([]string, error) {
	directories, err := listDirectories(path)
		if err != nil {
			return nil, err
		}

	var gitRepositories []string
	for _, dir := range directories {
		fullPath := path + "/" + dir
		if isGitRepository(fullPath) {
			gitRepositories = append(gitRepositories, dir)
		}
	}
	return gitRepositories, nil
}
