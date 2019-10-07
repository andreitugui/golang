package main

import "fmt"

func main() {
	slice1 := []string{"a", "b", "c"}
	slice2 := []string{"b", "a", "c"}
	l := len(slice1)
	var results []bool
	for i := 0; i <= l-1; i++ {
		results = append(results, check(slice1[i], slice2))
	}
	if checkResults(false, results) == true {
		fmt.Printf("the strings %v and %v are not anagrams", slice1, slice2)
	} else {
		fmt.Printf("the strings %v and %v are anagrams", slice1, slice2)
	}
}
func check(i string, s []string) bool {
	for _, a := range s {
		if a == i {
			return true
		}
	}
	return false
}
func checkResults(false bool, s []bool) bool {
	for _, a := range s {
		if a == false {
			return true
		}
	}
	return false
}
