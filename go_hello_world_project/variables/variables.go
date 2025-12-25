package main

import "fmt"

func main() {

	// In Go Lang, if we declare a variable with name then we have to use it otherwise it will be a compile time error.
	// var name string = "name"

	// By using this type, Go automatically infers it to relevant type.
	var name = "Ahmad & Maira"

	var isAdult = true

	var number = 20

	// Short Hand Declaration
	// It can only be used inside function outside function it will give some error.
	name2 := "Amma & Abba"

	var name3 string

	name3 = "Mama & Baba"

	fmt.Println(name)
	fmt.Println(isAdult)
	fmt.Println(number)

	fmt.Println(name2)
	fmt.Println(name3)
}
