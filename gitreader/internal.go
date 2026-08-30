package gitreader

import (
	"compress/zlib"
	"io"
	"log"
	"os"
	"strings"
)

// Return value of the HEAD of the repo
func getHead(repo string) (string, error) {
	fp, err := os.OpenFile(repo+"/.git/HEAD", os.O_RDONLY, 0)
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


// Get the commit object based on HEAD (not detached)
func getLatestCommit(repo string, head string) (string, error) {
	branchCommitObject := head[16 : len(head)-1]

	// Read the latest commit hash from the branch file
	branchFile := repo + "/.git/refs/heads/" + branchCommitObject
	fp, err := os.OpenFile(branchFile, os.O_RDONLY, 0)
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
		println("Error opening commit file:", err.Error())
		os.Exit(1)
	}
	defer fp.Close()

	// Decompress the data using zlib
	reader, err := zlib.NewReader(fp)
	if err != nil {
		println("Error creating zlib reader:", err.Error())
		os.Exit(1)
	}
	defer reader.Close()
	
	decompressedData, err := io.ReadAll(reader)
	if err != nil {
		println("Error reading decompressed data:", err.Error())
		os.Exit(1)
	}

	// Print the decompressed data
	log.Printf("Decompressed commit object data: %s", string(decompressedData))
	return string(decompressedData), nil
}