package main

import (
	"fmt"
	"sync"
)

type post struct {
	views int
	mu    sync.Mutex // Mutex can be added globally but best practice is to add it to the struct where it is needed.
}

func (p *post) incrementViews(wg *sync.WaitGroup) {
	//Another best practice is to use defer for unlocking the mutex so that even if there is a panic or error the mutex will be unlocked.
	defer func() {
		p.mu.Unlock() // Unlocking the critical section so that other goroutines can access it.
		wg.Done()
	}()
	p.mu.Lock() // Locking the critical section so that only one goroutine can access it at a time.
	p.views++   // Now this line is being accessed and modified by multiple lightweight threads(goroutines) simultaneously hence creating race condition. For this issue we will use mutex.
	// p.mu.Unlock() // Unlocking the critical section so that other goroutines can access it.

	// Mutex is good but has downside as it blocks the other goroutines until the mutex is unlocked. This can lead to performance issues if the critical section is large or if there are many goroutines trying to access it simultaneously. So better lock only the modification part.
}

// Using mutex for multithreading in Go. It basically used for preventing race condition.
func main() {
	wg := sync.WaitGroup{}

	myPost := post{
		views: 0,
	}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go myPost.incrementViews(&wg)
	}

	wg.Wait()
	fmt.Println(myPost.views)
}
