package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

func main() {
	text := "10 Write a program wordfreq to report the frequency of each word in an input text file. Call input.Split(bufio.ScanWords) before the first call to Scan to break the input into words instead of lines"
	words := strings.Fields(text)
	var freq map[string]int
	freq = make(map[string]int)
	for _, i := range words {
		fmt.Println(i)
		freq[i]++
	}
	b, _ := json.MarshalIndent(freq, "", "  ")
	fmt.Printf(string(b))
}
