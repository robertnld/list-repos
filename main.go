package main

import (
	"flag"
	"fmt"
)

func init() {
	// Define command-line flags
	flag.String("dir", ".", "Directory to list repositories from")
}

func main() {
	// Handle command-line flags
	flag.Parse()
	listDir := flag.Lookup("dir").Value.String()

	// Get the list of directories in the specified path
	directories, err := listDirectories(listDir)
	if err != nil {
		fmt.Println("Error listing directories:", err)
		return
	}

	// Print the list of directories in a web-interface
	for _, dir := range directories {

		fmt.Println(dir)
	}
}

