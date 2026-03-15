package main

import "fmt"

func main() {
	// var point = 8
	//
	// if point == 10 {
	// 	fmt.Println("lulus dengan nilai sempuarna")
	// } else if point > 5 {
	// 	fmt.Println("lulus")
	// } else if point == 4 {
	// 	fmt.Println("hampir lulus")
	// } else {
	// 	fmt.Printf("tidak lulus. nilai anda %d\n", point)
	// }

	// variable temporary

	// var point = 10000.0
	//
	// if percent := point / 100; percent >= 100 {
	// 	fmt.Printf("%.1f%s perceft!\n", percent, "%")
	// } else if percent >= 70 {
	// 	fmt.Printf("%.1f%s good\n", percent, "%")
	// } else {
	// 	fmt.Printf("%.1f%s not bad\n", percent, "%")
	// }

	// switch

	var point = 6

	switch point {
	case 8:
		fmt.Println("perfect")
	case 7:
		fmt.Print("awesome")
	default:
		fmt.Println("not bad")
	}

}
