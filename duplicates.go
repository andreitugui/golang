package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var content []byte
var splitLines []string
var files []string
var lines []string

func main() {
	files := os.Args[1:]
	for _, f := range files {
		content, _ = ioutil.ReadFile(f)
	}
	splitLines = strings.Split(string(content), "\n")
	lines = append(lines, splitLines...)
	// fmt.Printf("%s", lines)
	fmt.Printf("content:\n%s", content)
	//fmt.Printf("splitLines: %s", splitLines)
	// fmt.Printf("lines: %v", lines)
	var newlines []string
	i := 0
	for _, l := range lines {
		//		if contains(newlines, l) == false {
		if 1 == 1 {
			newlines = append(newlines, l)
		} else {
			i = i + 1
		}
	}
	// fmt.Printf("\nDuplicates: %v\n", i)
	// fmt.Printf("newlines: %v", newlines)
}
func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
