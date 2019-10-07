package main

import "fmt"

func main() {
	var s []string
	p := "aba"
	p2 := "fdsfds"
	push(p, s)
	push(p2, s)
	fmt.Println(s)
}
func push(p string, s []string) []string {
	s = append(s, p)
	return s
}
