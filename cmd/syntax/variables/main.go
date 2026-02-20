package main

import (
	"essentials/pkg/util"
	"fmt"
)

func main() {
	// manifest typing
	util.PrintHeader("Variabel")

	var firstName string = "John"
	var lastName string
	lastName = "Lennon"

	fmt.Printf("Halo %s %s!\n", firstName, lastName)

	//type interface

	middleName := "Widodo"

	fmt.Printf("Halo %s %s %s!\n", firstName, middleName, lastName)

	//multi variabel

	var first, second, third string
	first, second, third = "satu", "dua", "tiga"
	fmt.Println(first, second, third)

	var fourth, fifth, sixth string = "empat", "lima", "enam"
	fmt.Println(fourth, fifth, sixth)

	seventh, eight, ninth := "tujuh", "delapa", "sembila"
	fmt.Println(seventh, eight, ninth)

	one, isFriday, twoPointTwo, say := 1, true, 2.2, "hello"
	fmt.Println(one, isFriday, twoPointTwo, say)

	// reserved variabel

	_ = "ga dipake"
	name, _ := "john", "Widodo"
	fmt.Println(name)

	// variable new
	address := new(string)
	fmt.Println(address)
	fmt.Println(*address)

	// constant
	const satu = 1

	const (
		square         = "kotak"
		isToday bool   = true
		numeric uint16 = 120
	)

	const (
		today string = "monday"
		now
		isToday2 = true
	)

}
