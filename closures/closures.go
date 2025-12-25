package main

import "fmt"

// Closure example: A function that returns another function which increments and returns a counter.
func counter() func() int {
	count := 0

	return func() int {
		count++ // As count is defined in the outer function, so the value of count is preserved between calls.
		return count
	}
}

func main() {
	newCounter := counter()
	fmt.Println(newCounter()) // Output: 1
	fmt.Println(newCounter())
}
