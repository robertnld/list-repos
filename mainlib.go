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