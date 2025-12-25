package main

import "fmt"

func main() {
	// zeroed values
	//int -> 0, bool -> false, string -> ""

	// var numbs [4]int

	// Getting array length
	// fmt.Println(len(numbs))

	// numbs[0] = 1

	// fmt.Println(numbs)

	// var vals [4]bool
	// vals[2] = true
	// fmt.Println(vals)

	// var names [3]string
	// names[2] = "GoLang"
	// fmt.Println(names)

	// To declare it in single line
	// nums := [3]int{1, 2, 3}
	// fmt.Println(nums)

	//2d Arrays
	// nums := [2][2]int{{3, 4}, {5, 6}}
	// fmt.Println(nums)

	// - Array used for fixed size , that is predictable
	// - Memory optimization
	// - Constant time access

	var nums = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var target = 8
	fmt.Println(binarySearch(nums, target))
}

func binarySearch(nums [10]int, target int) int {

	var start = 0
	var end = len(nums) - 1

	for start <= end {
		var mid = start + (end-start)/2
		if nums[mid] == target {
			return mid
		} else if target < nums[mid] {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return 0
}
