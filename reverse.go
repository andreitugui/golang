package main

import "fmt"

func main() {
	sentence := "The amazing brown 狐 jumped over the small 犬"
	fmt.Println(len(sentence))
	runes := []rune(sentence)
	for i := len(runes) - 1; i >= 0; i-- {
		fmt.Printf("%v", string(runes[i]))
	}
}
