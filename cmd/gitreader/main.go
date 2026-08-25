package main

import (
	"list-repos/gitreader"
)


func main() {
	line, err := gitreader.Banner()
	if err != nil {
		println("Error:", err.Error())
		return
	}
	println(line)
}
