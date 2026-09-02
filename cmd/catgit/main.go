/*
Transform any Git object into readable format.
This is a simple implementation of the `git cat-file` command.

Usage:

	catgit [flags] [path ...]

The flag is:

	    -file string
			Git file to read

The app does not accept STDIN.
*/
package main

import (
	"flag"
	"fmt"
	"list-repos/gitreader"
	"os"
)

// Main function
// CLI only, no STDIN
func main() {
	filename := flag.String("file", "", "Git file to read")
	flag.Parse()

	// Must have parameter
	if flag.NFlag() == 0 {
		println("Usage: catgit -file <git_file>")
		os.Exit(1)
	}

	// Get the content of the Git file
	content, err := gitreader.ReadGitFile(*filename)
	if err != nil {
		fmt.Printf("Error reading Git file: %v\n", err)
		os.Exit(1)
	}

	// Print the content to STDOUT
	fmt.Println(content)
}