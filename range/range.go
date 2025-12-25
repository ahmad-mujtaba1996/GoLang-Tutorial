package main

import "fmt"

// iterating over data structures
func main() {
	nums := []int{10, 20, 30, 40, 50}

	sum := 0

	// Range over slice
	for i, num := range nums { // i is the index, num is the value
		sum += num
		fmt.Println(sum, i)
	}

	// Range over map
	m := map[string]string{"fname": "John", "lname": "Doe"}

	for k, v := range m { // k is the key, v is the value
		fmt.Println(k, v)
	}

	for k := range m { // only keys
		fmt.Println("Key:", k)
	}

	for _, v := range m { // only values
		fmt.Println("Value:", v)
	}

	//Range over string
	str := "Hello"
	for i, ch := range str {
		// i is the starting byte index, ch is the rune (Unicode code point)
		fmt.Println(i, string(ch)) // if tried to print ch directly, it would print the Unicode code point rune.
	}

}
