package main

import (
	"fmt"
	"sync"
)

// The main function where the execution starts
func main() {
	// Declare a WaitGroup to synchronize multiple goroutines
	var wg sync.WaitGroup
	// Launching multiple goroutines using a loop
	for i := 1; i <= 5; i++ {
		// Increment the WaitGroup counter before launching a goroutine
		wg.Add(1)

		// Launch a goroutine that executes myfunc
		go myfunc(&wg, i) // here pass by reference is manadatory otherwise will enter into a deadlock
	}

	// Wait for all goroutines to finish execution
	wg.Wait()

	// This line will execute only after all goroutines have finished
	fmt.Println("Each goroutine has run to completion, thanks for waiting!")
}

// Function executed by each goroutine
func myfunc(wg *sync.WaitGroup, i int) {
	// Ensure Done is called to decrement the WaitGroup counter once the function exits
	defer wg.Done() // if not done, it will enter a deadlock condition with error - all goroutines are asleep - deadlock!

	// Print a message indicating which iteration is executing
	fmt.Println("Finished executing iteration", i)
}
