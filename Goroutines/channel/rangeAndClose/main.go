package main

import (
	"fmt"
	"strconv"
	"sync"
)

func main() {
	// Create a buffered channel with a capacity of 2.
	// This means the channel can hold up to 2 messages before blocking further sends.
	var wg sync.WaitGroup
	ch := make(chan string, 2)

	// Start a goroutine to send messages.
	go func(ch chan string) {
		defer close(ch) // Ensure the channel is closed when the goroutine exits.
		wg.Add(1)
		for i := 1; i <= 1000; i++ { // Loop to send 5 messages.
			msg := "message" + strconv.Itoa(i)      // Construct a message.
			ch <- msg                               // Send the message into the channel.
			fmt.Printf("SEND goroutine: %v\n", msg) // Print message sent.
		}
	}(ch) // Pass the channel to the goroutine.

	// Start a goroutine to receive messages.
	go func(ch chan string) {
		defer wg.Done()
		// Use 'range' to continuously receive messages until the channel is closed.
		for val := range ch {
			fmt.Printf("RECV goroutine: %v\n", val) // Print received message.
		}

	}(ch) // Pass the channel to the goroutine.

	fmt.Println("MAIN goroutine: sleeping!") // Main goroutine pauses.

	// Allow some time for the send/receive goroutines to execute.
	// time.Sleep(time.Millisecond * 100)
	wg.Wait()

	// Attempt to receive from the channel after the sleep.
	val, ok := <-ch // 'ok' will be false if the channel is closed.
	fmt.Printf("Values returned from the read operation 'val, ok := <-ch' are: %q, %v\n", val, ok)

	fmt.Println("MAIN goroutine: done!") // Indicate main function is finished.
}

// output -
// MAIN goroutine: sleeping!
// SEND goroutine: message1
// SEND goroutine: message2
// RECV goroutine: message1
// RECV goroutine: message2
// SEND goroutine: message3
// RECV goroutine: message3
// SEND goroutine: message4
// RECV goroutine: message4
// SEND goroutine: message5
// RECV goroutine: message5
// Values returned from the read operation 'val, ok := <-ch' are: "", false
// MAIN goroutine: done!

// Key Observations
// Buffered Channel Behavior

// The buffer allows two messages to be stored without blocking.
// The third send blocks until a message is received.
// Goroutine Execution Order

// The SEND and RECV goroutines run concurrently.
// The first two sends happen immediately since the buffer can hold two messages.
// The third send waits until a message is received.
// Closing the Channel (defer close(ch))

// Ensures that the receiver exits gracefully after all messages are processed.
// The range loop in the receiver automatically stops when the channel is closed.
