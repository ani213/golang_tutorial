package main

import (
	"fmt"
	"sync"
)

func main() {
	// Slice of integers to process
	vals := []int{100, 50, 20, 90}

	// Generator function to send values into a channel
	in := gen(vals)

	// Fan-out: Two goroutines receive data from the 'in' channel and compute squares
	fo1 := square(in)
	fo2 := square(in)

	// Fan-in: Merge results from two channels into a single output channel
	fi := merge(fo1, fo2)

	// Receive and print results from the merged channel
	for res := range fi {
		fmt.Println(res)
	}
}

// gen function generates a channel and sends values from the given slice
func gen(vals []int) <-chan int {
	out := make(chan int) // Create an output channel

	go func() {
		// Send each value in 'vals' into the channel
		for _, val := range vals {
			out <- val
		}
		close(out) // Close channel after all values are sent
	}()

	return out // Return the output channel
}

// square function receives integers, computes squares, and sends them to another channel
func square(in <-chan int) <-chan int {
	out := make(chan int) // Create output channel for squared values

	go func() {
		for val := range in { // Read from input channel
			out <- val * val // Send squared value to output channel
		}
		close(out) // Close channel after processing all values
	}()

	return out // Return output channel
}

// merge function combines multiple input channels into a single output channel
func merge(fo ...<-chan int) <-chan int {
	out := make(chan int) // Create merged output channel
	var wg sync.WaitGroup // WaitGroup to synchronize goroutines

	// Function to read from a channel and send values to the output channel
	fi := func(ch <-chan int) {
		for val := range ch { // Read from input channel
			out <- val // Send value to output channel
		}
		wg.Done() // Mark this goroutine as done
	}

	// Start a goroutine for each input channel
	wg.Add(len(fo)) // Add to WaitGroup based on the number of input channels
	for _, ch := range fo {
		go fi(ch) // Launch a goroutine to read from each channel
	}

	// Close the output channel after all input channels are processed
	go func() {
		wg.Wait()  // Wait for all goroutines to complete
		close(out) // Close the output channel
	}()

	return out // Return merged output channel
}
