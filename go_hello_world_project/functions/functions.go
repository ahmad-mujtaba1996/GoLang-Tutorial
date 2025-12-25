package main

import "fmt"

// func add(a int, b int) (int, int) {
// 	return a + b, a - b
// }

// if function parameters have same type, we can omit the type for all but the last one
// func add(a, b int) (int, int) {
// 	return a + b, a - b
// }

// return tuples
// func add(a, b int) (sum int, diff int) {
// 	sum = a + b
// 	diff = a - b
// 	return sum, diff
// }

// function returns an array or slice
// func getNames() []string {
// 	return []string{"Ahmad", "Ali", "Omar"}
// }

// function that returns a map
// func getLanguages() map[int]string {
// 	return map[int]string{0: "Go", 1: "Python", 2: "JavaScript"}
// }

// Passing function as parameter
// func processIt(fn func(a int) int) {
// 	value := 10
// 	result := fn(value)
// 	fmt.Println("Processed Result is:", result)
// }

// returning function from function
func processIt() func(a int) int {
	return func(a int) int {
		fmt.Println("Second Call")
		fmt.Println("Value is:", a)
		return a * 2
	}
}

func main() {
	// fn := func(a int) int {
	// 	fmt.Println("Value is:", a)
	// 	return a * 2
	// }
	// processIt(fn)

	fn := processIt()
	fmt.Println("First Call")
	result := fn(15)
	fmt.Println("Processed Result is:", result)

	// result, diff := add(5, 3)
	// fmt.Println("Resultant Value is:", result, "Difference is:", diff)
}
