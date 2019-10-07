package main

import "fmt"

func main() {
	var zero int
	var negative int
	var positive int
	array := [5]int{1, 1, 0, -1, -1}
	for _, el := range array {
		switch {
		case el == 0:
			zero = zero + 1
		case el < 0:
			negative = negative + 1
		case el > 0:
			positive = positive + 1
		}
	}
	zeroFloat := float64(zero) / float64(len(array))
	negativeFloat := float64(negative) / float64(len(array))
	positiveFloat := float64(positive) / float64(len(array))
	fmt.Printf("%f\n", zeroFloat)
	fmt.Printf("%f\n", negativeFloat)
	fmt.Printf("%f", positiveFloat)
}
