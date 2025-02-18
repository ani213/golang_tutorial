package main

import (
	"fmt"
	"time"
)

func main() {
	// Creating three unbuffered channels for communication between goroutines
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	// Goroutine to send a value into ch1 after 1 second
	go func() {
		time.Sleep(1 * time.Second) // Simulating delay
		// time.Sleep(10 * time.Second) // Uncomment to simulate a longer delay
		ch1 <- 1 // Sending value 1 into ch1
	}()

	// Goroutine to send a value into ch2 after 2 seconds
	go func() {
		time.Sleep(2 * time.Second) // Simulating delay
		// time.Sleep(10 * time.Second) // Uncomment to simulate a longer delay
		ch2 <- 2 // Sending value 2 into ch2
	}()

	// Goroutine to send a value into ch3 after 3 seconds
	go func() {
		time.Sleep(3 * time.Second) // Simulating delay
		// time.Sleep(10 * time.Second) // Uncomment to simulate a longer delay
		ch3 <- 3 // Sending value 3 into ch3
	}()

	// Using select to receive values from any of the three channels
	for i := 0; i < 3; i++ { // Loop to receive three values
		select {
		case val1 := <-ch1: // If ch1 sends a value, this case executes
			fmt.Printf("value received from ch1: %v\n", val1)
		case val2 := <-ch2: // If ch2 sends a value, this case executes
			fmt.Printf("value received from ch2: %v\n", val2)
		case val3 := <-ch3: // If ch3 sends a value, this case executes
			fmt.Printf("value received from ch3: %v\n", val3)
			// case to := <-time.After(5 * time.Second):
			// 	fmt.Printf("timed out after 5 seconds at %v\n", to)
			// 	return // Exit if timeout occurs
		}
	}
}

// output -
// value received from ch1: 1
// value received from ch2: 2
// value received from ch3: 3

// Explanation:
// Channel Creation:

// Three unbuffered channels ch1, ch2, and ch3 are created. They will be used to receive integer values from goroutines.
// Goroutines:

// Three separate goroutines are started. Each one sleeps for a specific duration (1s, 2s, and 3s) before sending a value to its respective channel.
// Receiving Data Using select:

// A for loop runs three times to ensure all three values are received.
// The select statement listens to multiple channels concurrently. The first channel that sends data is processed.
// Whichever channel sends a value first will have its case executed.
// Timeout Mechanism (Commented Out):

// The time.After(5 * time.Second) case (commented) provides a way to break the waiting if no value is received within 5 seconds.
// If uncommented, this case would execute if no channels send data within the timeout duration.
