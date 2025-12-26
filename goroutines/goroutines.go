package main

import (
	"fmt"
	"sync"
)

func task(id int, wg *sync.WaitGroup) {
	defer wg.Done() // This is required to notify wait group that this goroutine is completed. Using defer keyword, we make sure that this function is called once all other functionalities in function are completedeven if there is any error in the function.
	fmt.Println("Task", id)
}

func main() {

	var wg sync.WaitGroup // Using wait group, we can hold main functions to wait for all goroutines to complete. It still doesn't check hierarchy as our goroutines run in parallel.

	for i := 0; i <= 10; i++ {
		// task(i) // By this way, we run multiple tasks by blocking and running in a hierarchical order, which will be slower compared to goroutines
		// go task(i)// By this way, we run multiple tasks by running them in parallel, which will be faster and its order will be different compared to hierarchical order

		// Inline function
		// func() {
		// 	fmt.Println("Task", i)
		// }()

		// Inline function with go
		// go func(index int) {
		// 	fmt.Println("Task", index)
		// }(i)

		wg.Add(1)
		go task(i, &wg)
	}

	wg.Wait()
	fmt.Println("All Tasks are completed")
}
