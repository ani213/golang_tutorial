package bufferedVersionOfUnbuffered

import (
	"fmt"
	"time"
)

func BufferedVersionOfUnbuffered() {
	// Create a buffered channel with a capacity of 1
	ch := make(chan string, 1)

	// Start the 'send' goroutine, which will attempt to send a message into the channel.
	go send1(ch)

	// Print the length of the channel after sending (should be 0 initially but it can be increase to its capacity).
	fmt.Printf("channel length: %v\n", len(ch))

	// Start the 'receive' goroutine to receive the message after a delay.
	go receive1(ch)

	// Sleep to allow the goroutines to complete execution before 'main' exits.
	time.Sleep(time.Second * 2)
}

// Function to send a message into the channel.
func send1(ch chan string) {
	fmt.Println("blocking send goroutine...")

	// Sending a message into an buffered channel.
	// Since the channel is buffered, this line will not block. This line executes immediately because the message fits in the buffer..
	ch <- "message"

	// This line executes immediately.
	fmt.Println("send goroutine unblocked...")
}

// Function to receive a message from the channel.
func receive1(ch chan string) {
	// Simulate some delay before receiving the message.
	time.Sleep(time.Second * 1)

	// Print the channel length (should still be 1, since it's an buffered channel).
	fmt.Printf("channel length: %v\n", len(ch))

	// Receive the message from the channel, which will unblock the 'send' function(if it is blocked).
	fmt.Printf("received: %v\n", <-ch)
}

// output -
// channel length: 0
// blocking send goroutine...
// send goroutine unblocked...
// channel length: 1
// received: message

// Unlike an unbuffered channel, this can hold 1 value without blocking.
// Since the buffer has space, the send operation does not block and send continues execution.
// Since the message is stored in the buffer, len(ch) == 1.
