package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	// Set the number of CPU cores to be used
	runtime.GOMAXPROCS(runtime.NumCPU())

	// Shared counter variable
	var counter int64

	// WaitGroup to synchronize goroutines
	var wg sync.WaitGroup

	// Mutex for safe concurrent access to counter
	var mu sync.Mutex

	// Function to safely increment the counter
	inc := func() {
		mu.Lock()         // Lock to prevent race conditions
		defer mu.Unlock() // Ensure unlock even if function exits early
		counter++         // Increment counter
		fmt.Println("Incremented counter =", counter)
	}

	// Function to safely decrement the counter
	dec := func() {
		mu.Lock()         // Lock to prevent race conditions
		defer mu.Unlock() // Ensure unlock even if function exits early
		counter--         // Decrement counter
		fmt.Println("Decremented counter =", counter)
	}

	// more optimised approach is to use atomic

	// inc := func() {
	// 	atomic.AddInt64(&counter, 1)
	// 	fmt.Println("Incremented counter =", counter)
	// }

	// dec := func() {
	// 	atomic.AddInt64(&counter, -1)
	// 	fmt.Println("Decremented counter =", counter)
	// }

	// Launch 1000 goroutines for incrementing the counter
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			inc()
		}()
	}

	// Launch 1000 goroutines for decrementing the counter
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			dec()
		}()
	}

	// Wait for all goroutines to finish execution
	wg.Wait()

	// Print the final counter value (should be 0)
	fmt.Println("Final value of counter:", counter)
}
