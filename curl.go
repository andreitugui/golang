package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	urls := os.Args[1:]

	for _, url := range urls {
		resp, _ := http.Get(url)
		fmt.Printf("%v", resp)
		fmt.Println("\n")
	}
}
