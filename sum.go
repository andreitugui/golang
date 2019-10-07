package main

import "fmt"

func main() {
	var sum int
	numbers := [5]int{1, 3, 5, 6, 1117}
	for _, no := range numbers {
		sum = sum + no
	}
	fmt.Println(sum)
}
