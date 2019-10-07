package main

import "fmt"

func main() {
	var l1, v1, l2, v2 int = 0, 2, 5, 3
	var s int
	s = (l1 - l2) / (v2 - v1)
	if s > 0 {
		fmt.Printf("yes, after %v jumps", s)
	} else {
		fmt.Println("no")
	}
}
