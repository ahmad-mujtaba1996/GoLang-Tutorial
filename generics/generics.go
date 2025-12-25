package main

import "fmt"

// Generic use with structs.
// type Stack[T any] struct {
// 	elements []T
// }

// Generic function with parameters as generic

// using any keyword is considered a bad practice because we can pass any kind of values referring to its datatype
// func printSlice[T any](items []T) {
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }

// Apart from any, we can use interface{} as well. It will work in a similar manner
// func printSlice[T interface{}](items []T) {
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }

// If we want to restrict the generic function to some type then using pipe operator we can do it.
// func printSlice[T int | string | bool](items []T) {
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }

// Using comparable operator we can also make our function generic
// func printSlice[T comparable](items []T) {
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }

// Using multiple generics
func printSlice[T comparable, V string](items []T, name V) {
	for _, item := range items {
		fmt.Println(item, name)
	}
}
func main() {
	// nums := []int{1, 2, 3, 4, 5}
	// names := []string{"John", "Doe", "Ahmad"}
	// flags := []bool{true, false, true}
	// printSlice(nums)
	// printSlice(names)
	// printSlice(flags) // Its will give error if bool is not allowed in print slice generic

	// values := []int{1, 2, 3, 4, 5}
	// stack := Stack[int]{ // here we are passing int as generic type to make sure our generic stack accepts int only values
	// 	elements: values,
	// }
	// stringStack := Stack[string]{ // here we are passing int as generic type to make sure our generic stack accepts int only values
	// 	elements: []string{"a", "b", "c"},
	// }
	// fmt.Println(stack)
	// fmt.Println(stringStack)

	//Using multiple generics
	printSlice([]int{1, 2, 3, 4, 5}, "John")
}
