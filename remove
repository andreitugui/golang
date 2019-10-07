package main

import "fmt"

func main() {
	array := []string{"0", "1", "2", "3", "4", "5"}
	remove := 2
	l := len(array)
	var newArray []string
	for i := 0; i < remove; i++ {
		newArray = append(newArray, array[i])
	}
	for i := remove + 1; i < l; i++ {
		newArray = append(newArray, array[i])
	}
	fmt.Println(newArray)
}
