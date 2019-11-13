package main

import "fmt"

func main() {
	matrix := [3][3]int{
		{1, 2, 5},
		{4, 5, 6},
		{9, 8, 9},
	}
	var diag1 int
	var diag2 int
	var i int
	l := len(matrix)
	for i = 0; i <= 2; i++ {
		diag1 = diag1 + matrix[i][i]
	}
	for i = 0; i <= 2; i++ {
		diag2 = diag2 + matrix[i][l-i-1]
	}
	var diff int
	if diag1 > diag2 {
		diff = diag1 - diag2
	}
	diff = diag2 - diag1
	fmt.Println(diff)
}
