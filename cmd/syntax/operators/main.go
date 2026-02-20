package main

import (
	"essentials/pkg/util"
	"fmt"
)

func main() {
	util.PrintHeader("Operator")

	var value = (((2+6)%3)*4 - 2) / 3
	fmt.Printf("hasil %d", value)

	var left = false
	var right = true

	var leftAndRight = left && right
	fmt.Printf("left && right \t(%t) \n", leftAndRight)

	var leftOrRight = left || right
	fmt.Printf("left || right \t(%t) \n", leftOrRight)

	var leftReverse = !left
	fmt.Printf("!left \t\t(%t) \n", leftReverse)
}

