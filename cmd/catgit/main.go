/*
Transform any Git object into readable format.
This is a simple implementation of the `git cat-file` command.

Usage:

	catgit --file <path>

The flag is:

	    -file string
			Git file to read

The app does not accept STDIN.
*/
package main

import (
	"flag"
	"fmt"
	"io"
	"list-repos/gitreader"
	"os"
)

// CLI only, no STDIN
func main() {
	filename := flag.String("file", "", "Git file to read")
	flag.Parse()

	// Must have parameter
	if *filename == "" {
		fmt.Fprintf(os.Stderr, "Usage: catgit -file <git_file>\n")
		os.Exit(1)
	}

	// Get the content of the Git file
	content, err := gitreader.ReadGitFile(*filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading Git file: %v\n", err)
		os.Exit(1)
	}

	_, err = io.WriteString(os.Stdout, content)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error writing to STDOUT: %v\n", err)
		os.Exit(1)
	}
}
