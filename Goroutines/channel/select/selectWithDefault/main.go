package main

import (
	"fmt"
	"time"
)

func main() {
	// Creating two unbuffered channels for communication
	ch1 := make(chan int)
	ch2 := make(chan int)

	// Goroutine to send data into channels
	go func() {
		for i := 1; i < 4; i++ { // Loop runs 3 times
			time.Sleep(3 * time.Second) // Wait for 3 seconds before sending to ch1
			ch1 <- i                    // Send value i to ch1

			time.Sleep(1 * time.Second) // Wait for 1 second before sending to ch2
			ch2 <- i                    // Send value i to ch2
		}
	}()

	// Loop runs 4 times, attempting to receive from channels
	for i := 1; i < 5; i++ {
		select {
		case val := <-ch1: // If ch1 sends a value, this case executes
			fmt.Printf("value received from channel ch1: %v\n", val)
		case val := <-ch2: // If ch2 sends a value, this case executes
			fmt.Printf("value received from channel ch2: %v\n", val)
		default: // If no data is available, execute this block
			fmt.Println("no data received... performing some other operation")
		}
		time.Sleep(2 * time.Second) // Simulate processing time before next iteration
	}
}

// output -
// no data received... performing some other operation
// no data received... performing some other operation
// value received from channel ch1: 1
// value received from channel ch2: 1

// Explanation:
// Channel Creation:

// ch1 and ch2 are unbuffered channels used for integer communication.
// Goroutine to Send Data:

// A loop runs three times, sending values to ch1 after 3 seconds and to ch2 after an additional 1 second.
// Receiving Data Using select:

// The loop runs four times.
// The select statement listens for data from ch1 and ch2.
// If neither channel is ready, the default case executes, indicating no data was received.
// After each iteration, a 2-second delay simulates processing time.
