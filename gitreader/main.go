package gitreader

import (
	"log"
	"os"
	"os/exec"
)

// PrintBanner returns a banner message as test
func Banner() (string, error) {
	msg := `List-repos for listing Git repos ...`
	return msg, nil
}

// Test that directory is a Git repository by checking for the presence of a .git folder
func IsGitRepository(path string) bool {
	gitDir := path + "/.git"
	info, err := os.Stat(gitDir)
	if err != nil {
		return false
	}
	return info.IsDir()
}


// Get HEAD of the repo
func GetHead(repo string) (string, error) {
	fp, err := os.OpenFile(repo + "/.git/HEAD", os.O_RDONLY, 0)
	if err != nil {
		println("Error opening file:", err.Error())
		os.Exit(1)
	}
	data, err := os.ReadFile(fp.Name())
	if err != nil {
		println("Error reading file:", err.Error())
		os.Exit(1)
	}
	defer fp.Close()
	log.Printf("HEAD: %s", string(data))
	return string(data), nil
}


// Latest commit message for the Git repository at the specified path
func GetLatestCommitMessage(path string) (string, error) {
	cmd := exec.Command("git", "-C", path, "log", "-1", "--pretty=%B")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(output), nil
}
