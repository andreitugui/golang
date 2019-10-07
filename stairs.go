package main

import "fmt"

func main() {
	n := 10
	for x := 0; x < n; x++ {
		for i := 1; i <= n-x-1; i++ {
			fmt.Printf(" ")
		}
		for i := 0; i <= x; i++ {
			fmt.Printf("#")
		}
		fmt.Printf("\n")
	}
}
