package main

// slice -> dynamic array
// Most used construct in Go
// + useful methods
func main() {
	// Decalared but not initialized
	// var numbs []int
	// fmt.Println(numbs == nil)
	// fmt.Println(len(numbs))

	// If you want to make sure you slice isn't set to nil
	// var numbs = make([]int, 0, 5) // -> 0 is capacity where by default values will be available like 0 but 5 is initial capacity incase if we know the actual length
	// capacity -> max numbers of element a slice can hold
	// fmt.Println(cap(numbs))
	// numbs = append(numbs, 1)
	// numbs = append(numbs, 2)
	// numbs = append(numbs, 3)
	// numbs = append(numbs, 4)
	// fmt.Println(numbs)
	// fmt.Println(cap(numbs))
	// fmt.Println(len(numbs))

	// numbs := []int{}
	// numbs = append(numbs, 1)
	// numbs = append(numbs, 2)
	// numbs = append(numbs, 3)
	// numbs = append(numbs, 4)
	// numbs = append(numbs, 5)
	// fmt.Println(numbs)
	// fmt.Println(cap(numbs))
	// fmt.Println(len(numbs))

	// In case we need to modify specific index value.
	// var numbs = make([]int, 2, 5) // -> If 0 is length then we will get index out of range error hence need to setup some length or append at start.
	// numbs[0] = 11
	// fmt.Println(numbs)
	// fmt.Println(cap(numbs))
	// fmt.Println(len(numbs))

	// copy function
	// var numbs = make([]int, 0, 5)
	// numbs = append(numbs, 2) // need to append something in parent slice in order to get it copied in new slice
	// var newNumbs = make([]int, len(numbs))
	// copy(newNumbs, numbs)
	// fmt.Println(numbs, newNumbs)

	// slice operator
	// var nums = []int{1, 2, 3}
	// fmt.Println(nums[0:2]) // slice operator excludes the to index(in this case, to index is "2") value that is 3 and prints other values.
	// fmt.Println(nums[0:1]) // slice operator excludes the to index(in this case, to index is "1") value that is 2 and prints other values.
	// fmt.Println(nums[:1]) // if from index is not provided it will consider from 0th index
	// fmt.Println(nums[1:]) // if to index is not provided it will consider till last index

	// slices
	// var nums = []int{1,2}
	// var nums2 = []int{1,2}
	// Comparing two slices
	// fmt.Println(slices.Equal(nums, nums2))

	// 2d slices
	// var nums = [][]int{{1, 2}, {3, 4}, {5, 6}}
	// fmt.Println(nums)
}
