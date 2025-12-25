package main

import "fmt"

// Pointers contains the memory address of a value
// by value means that the function receives a copy of the value
// func changeNum(num int) { // by value its copied
// 	num = 5 // this is a copy of the value passed so actuall value isn't changed.
// 	fmt.Println("In changeNum", num)
// }

// by reference means that the function receives the memory address of the value
func changeNum(num *int) {
	*num = 5 // * is used to dereference the pointer and get the value it points to
	fmt.Println("In changedNum", *num)

}

func main() {
	// num := 1
	// changeNum(num)
	// fmt.Println("After changeNum in main:", num)

	num := 1
	fmt.Println("Memory address1", &num) // In order to get the memory address of a value we use the & operator
	changeNum(&num)
	fmt.Println("After changeNum in main:", num)
}
