package main

import (
	"compress/zlib"
	"flag"
	"io"
	"os"
)

func main() {
	filename := flag.String("file", "", "Git file to read")
	flag.Parse()

	fp, err := os.OpenFile(*filename, os.O_RDONLY, 0)
	if err != nil {
		println("Error opening file:", err.Error())
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
	println(string(decompressedData))
}