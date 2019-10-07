package main

import "fmt"

// InputMiles 123
var InputMiles float64 = 10

func converter(InputMiles float64) float64 {
	const conv = 1.609344
	OutputKms := float64(InputMiles * conv)
	return OutputKms
}
func main() {
	print := converter(InputMiles)
	fmt.Printf("%.2f miles = %f km", InputMiles, print)
}
