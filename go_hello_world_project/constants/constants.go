package main

import "fmt"

const ageLimit = 10

func main() {

	//Const variable cannot be changed
	const name = "Ahmad"
	const age = 25

	fmt.Println(ageLimit)

	// Constant Grouping
	const (
		port = 5000
		host = "localhost"
	)

	fmt.Println(host, port)
}
