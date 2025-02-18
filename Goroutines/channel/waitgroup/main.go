package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup // Declare a WaitGroup to synchronize goroutines
	ch := make(chan int)  // Create an unbuffered channel for integer communication

	wg.Add(3) // We have 3 goroutines that will send data, so increment WaitGroup counter by 3

	// Goroutine to close the channel once all sending goroutines complete
	go func() {
		wg.Wait() // Wait for all 3 goroutines to finish
		close(ch) // Close the channel to signal that no more values will be sent
	}()

	// Goroutine 1: Sends the value 1 into the channel
	go func() {
		defer wg.Done() // Decrement the WaitGroup counter when done
		ch <- 1         // Send value 1 to the channel
	}()

	// Goroutine 2: Sends the value 2 into the channel
	go func() {
		defer wg.Done() // Decrement the WaitGroup counter when done
		ch <- 2         // Send value 2 to the channel
	}()

	// Goroutine 3: Sends the value 3 into the channel
	go func() {
		defer wg.Done() // Decrement the WaitGroup counter when done
		ch <- 3         // Send value 3 to the channel
	}()

	// Main goroutine: Receiving values from the channel using range
	for val := range ch {
		fmt.Printf("Value received: %v\n", val)
	}

	// Attempting to receive from a closed channel
	// Once a channel is closed, any further receive operations return the zero-value for the channel’s type
	// and 'ok' will be false
	val, ok := <-ch
	fmt.Println(val, ok) // Output: 0 false
}

// output -
// Value received: 1
// Value received: 2
// Value received: 3
// 0 false

// 1. Synchronization Using sync.WaitGroup
// The sync.WaitGroup (wg) is used to make sure all goroutines complete before closing the channel.
// wg.Add(3) sets the counter to 3, as there are 3 goroutines sending data.
// Each sender goroutine calls wg.Done() once it sends a value, decrementing the counter.
// 2. Closing the Channel Safely
// A separate goroutine calls wg.Wait() to wait for all sender goroutines to finish.
// Only after all sender goroutines are done, it calls close(ch).
// This ensures the channel is never closed while it's still being written to, preventing a "send on closed channel" panic.
// 3. Using range to Receive from the Channel
// The for val := range ch loop reads values from ch until it's closed.
// Since the channel is properly closed, this loop terminates automatically after receiving all values.
// 4. Reading from a Closed Channel
// The final val, ok := <-ch statement attempts to read from ch after it has been closed.
// Since ch is closed:
// val returns 0 (default zero value of int).
// ok returns false, indicating the channel is closed.
