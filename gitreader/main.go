package gitreader

import (
	"os"
	"os/exec"
)


func PrintBanner() {
	println("Starting gitreader...")
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


// Latest commit message for the Git repository at the specified path
func GetLatestCommitMessage(path string) (string, error) {
	cmd := exec.Command("git", "-C", path, "log", "-1", "--pretty=%B")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(output), nil
}


func main () {
	println("Starting cli...")
}