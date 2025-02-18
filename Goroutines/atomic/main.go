package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

// The main function where execution begins
func main() {
	// Set the number of OS threads to the number of CPU cores
	runtime.GOMAXPROCS(runtime.NumCPU()) //Ensures that the program utilizes all available CPU cores for parallel execution.

	// Declare an unsigned 64-bit integer counter
	var counter uint64

	// Declare a WaitGroup to synchronize goroutines
	var wg sync.WaitGroup

	// Loop to spawn 1000 goroutines
	for i := 0; i < 1000; i++ {
		// Increment the WaitGroup counter before launching a goroutine
		wg.Add(1)

		// Launch a goroutine to increment the counter
		go func() {
			// Ensure Done() is called when the function exits
			defer wg.Done()

			// Loop to increment the counter 1000 times
			for j := 0; j < 1000; j++ {
				// counter++ // Unsafe increment due to concurrent access Since multiple goroutines modify counter concurrently,
				// the final result will likely be incorrect. counter++ is not an atomic operation, leading to a race condition.
				atomic.AddUint64(&counter, 1)
				//  ensures thread safety by preventing race conditions, as only one goroutine can access the counter variable at a time
			}
		}()
	}

	// Wait for all goroutines to complete execution
	wg.Wait()

	// Print the final value of the counter
	fmt.Println("counter:", counter)
}
