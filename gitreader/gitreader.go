package gitreader

import (
	"os"
)

// Test that directory is a Git repository by checking for the presence of a .git folder
func IsGitRepository(path string) bool {
	gitDir := path + "/.git"
	info, err := os.Stat(gitDir)
	if err != nil {
		return false
	}
	return info.IsDir()
}


// Latest commit message of the Git repository at the given path
func GetLatestCommitMessage(repo string) (string, error) {
	
	repoHead, err := getHead(repo)
	if err != nil {
		return "", err
	}

	repoLatestCommit, err := getLatestCommit(repo, repoHead)
	if err != nil {
		return "", err
	}

	commitObject, err := getCommitObjectPath(repo,repoLatestCommit)
	if err != nil {
		return "", err
	}

	commitMessage, err := getCommitMessage(commitObject)
	if err != nil {
		return "", err
	}

	return commitMessage, nil
} 