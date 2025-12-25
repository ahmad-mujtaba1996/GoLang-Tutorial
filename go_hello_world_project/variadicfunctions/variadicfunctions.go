package main

import "fmt"

func sum(nums ...int) int { // ...int indicates a variadic function
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func main() {
	fmt.Println(1, 2, 3, 4, 5, "Hello")

	result := sum(1, 2, 3, 4, 5)
	fmt.Println("Sum:", result)

	nums := []int{2, 4, 6, 8}
	result = sum(nums...) // passing nums to sum function using spread operator incase slice is defined explicitly
	fmt.Println("Sum of slice:", result)
}
