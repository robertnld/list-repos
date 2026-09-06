package gitreader

import (
	"bytes"
	"compress/zlib"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// Test that directory is a Git repository by checking for the presence of a .git folder
func IsGitRepository(path string) (bool, error) {
	// Convert the incoming directory path to an absolute path
	absPath, err := filepath.Abs(path)
	if err != nil {
		return false, fmt.Errorf("error getting absolute path: %v", err)
	}

	gitDir := absPath + "/.git"
	info, err := os.Stat(gitDir)
	if err != nil {
		if os.IsNotExist(err) {
			return false, fmt.Errorf(".git directory does not exist")
		} else {
			return false, fmt.Errorf("error stating .git directory: %v", err)
		}
		
	}
	if !info.IsDir() {
		return false, fmt.Errorf(".git exists but is not a directory")
	}

	return true, nil
}


// Read the contents of a Git file and return it as a string
// For now, decompress only and keep the header in the string
func ReadGitFile(path string) (string, error) {
	// Open the file for reading
	fp, err := os.OpenFile(path, os.O_RDONLY, 0)
	if err != nil {
		return "", fmt.Errorf("error opening file: %v", err)
	}
	defer fp.Close()

	// Create the required zlib reader
	reader, err := zlib.NewReader(fp)
	if err != nil {
		return "", fmt.Errorf("error creating zlib reader: %v", err)
	}
	defer reader.Close()
	
	// Decompress the data using zlib
	decompressedData, err := io.ReadAll(reader)
	if err != nil {
		return "", fmt.Errorf("error reading decompressed data: %v", err)
	}
	return string(decompressedData), nil
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

func TypeOfGitObject(data string) (string, error) {
	// Return first part
	byteData := []byte(data)
	i := bytes.IndexByte(byteData, 0)
	if i == -1 {
    	return "", fmt.Errorf("missing NUL byte")
	}

	header := string(data[:i])
	/* content := data[i+1:]

	parts := strings.SplitN(header, " ", 2)

	objectType := parts[0]
	size, err := strconv.Atoi(parts[1])
	if err != nil {
    	return "", err
	} */
	return header, nil
}

