package main

import (
	"essentials/pkg/util"
	"fmt"
)

func main() {
	util.PrintHeader("Data Type")
	var positiveNumber uint8 = 89
	var negativeNumber = -12273910281921223

	fmt.Printf("bilangan positif:  %d\n ", positiveNumber)
	fmt.Printf("bilangan negatif: %d\n ", negativeNumber)

	var decimalNumber = 2.62

	fmt.Printf("bilangan desimal: %f\n", decimalNumber)
	fmt.Printf("bilangan desimal: %.3f\n", decimalNumber)

	var exists bool = true
	fmt.Printf("exists ? %t\n", exists)
}
