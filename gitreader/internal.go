package gitreader

import (
	"compress/zlib"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

// Return value of the HEAD of the repo
func getHead(repo string) (string, error) {
	fp, err := os.OpenFile(repo+"/.git/HEAD", os.O_RDONLY, 0)
	if err != nil {
		return "",fmt.Errorf("error opening file: %v", err)
	}
	data, err := os.ReadFile(fp.Name())
	if err != nil {
		return "", fmt.Errorf("error reading file: %v", err)
	}
	defer fp.Close()
	log.Printf("HEAD: %s", string(data))
	return string(data), nil
}


// Get the commit object based on HEAD (not detached)
func getLatestCommit(repo string, head string) (string, error) {
	branchCommitObject := head[16 : len(head)-1]

	// Read the latest commit hash from the branch file
	branchFile := repo + "/.git/refs/heads/" + branchCommitObject
	fp, err := os.OpenFile(branchFile, os.O_RDONLY, 0)
	if err != nil {
		return "", fmt.Errorf("error opening file: %v", err)
	}
	data, err := os.ReadFile(fp.Name())
	if err != nil {
		return "", fmt.Errorf("error reading file: %v", err)
	}
	defer fp.Close()
	log.Printf("Latest commit object hash: %s", string(data))
	return string(data), nil
}


// Get the path of the commit object for the given commit hash
func getCommitObjectPath(repo string, commitHash string) (string, error) {

	commitObjectPath := (
		strings.TrimSpace(
			repo + "/.git/objects/" + commitHash[:2] + "/" + commitHash[2:]))
	log.Printf("Commit object path: %s", commitObjectPath)
	return commitObjectPath, nil
}


// Get the commit message from the commit object file
func getCommitMessage(commitObjectPath string) (string, error) {

	fp, err := os.OpenFile(commitObjectPath, os.O_RDONLY, 0)
	if err != nil {
		return "", fmt.Errorf("error opening commit file: %v", err)
	}
	defer fp.Close()

	// Decompress the data using zlib
	reader, err := zlib.NewReader(fp)
	if err != nil {
		return "", fmt.Errorf("error creating zlib reader: %v", err)
	}
	defer reader.Close()
	
	decompressedData, err := io.ReadAll(reader)
	if err != nil {
		return "", fmt.Errorf("error reading decompressed data: %v", err)
	}

	// Print the decompressed data
	log.Printf("Decompressed commit object data: %s", string(decompressedData))
	return string(decompressedData), nil
}