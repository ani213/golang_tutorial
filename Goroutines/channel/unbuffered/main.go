package main

import (
	"fmt"
	"go-tutorials/Goroutines/channel/unbuffered/bufferedVersionOfUnbuffered"
	"go-tutorials/Goroutines/channel/unbuffered/limitedChannelLength"
	"time"
)

func main() {

	// Create an unbuffered channel for string messages.
	ch := make(chan string)

	// Start the 'send' goroutine, which will attempt to send a message into the channel.
	go send(ch)

	// Print the length of the channel (always 0 for an unbuffered channel unless a value is pending).
	fmt.Printf("channel length: %v\n", len(ch))

	// Start the 'receive' goroutine to receive the message after a delay.
	go receive(ch)
	time.Sleep(time.Second * 2)
	fmt.Print("\n-------BufferedVersionOfUnbuffered--------\n")
	bufferedVersionOfUnbuffered.BufferedVersionOfUnbuffered()

	fmt.Print("\n-------LimitedChannelLength--------\n")
	limitedChannelLength.LimitedChannelLength()
	// Sleep to allow the goroutines to complete execution before 'main' exits.

}

// Function to send a message into the channel.
func send(ch chan string) {
	fmt.Println("blocking send goroutine...")

	// Sending a message into an unbuffered channel.
	// Since the channel is unbuffered, this line will block until the 'receive' function consumes the message.
	ch <- "message"

	// This line executes only after the message is received.
	fmt.Println("send goroutine unblocked...")
}

// Function to receive a message from the channel.
func receive(ch chan string) {
	// Simulate some delay before receiving the message.
	time.Sleep(time.Second * 1)

	// Print the channel length (should still be 0, since it's an unbuffered channel).
	fmt.Printf("channel length: %v\n", len(ch))

	// Receive the message from the channel, which will unblock the 'send' function.
	fmt.Printf("received: %v\n", <-ch)
}

// output -
// blocking send goroutine...
// channel length: 0
// channel length: 0
// received: message
// send goroutine unblocked...

// An unbuffered channel means that sending to ch <- "message" will block until another goroutine reads from it.
// The send function attempts to send "message" into ch but gets blocked because no goroutine is ready to receive.
// Since the channel is unbuffered, its length is 0.
// The send goroutine remains blocked at ch <- "message".
// The receive function calls <-ch, reading "message" from the channel.
// This unblocks the send function, allowing it to proceed

// Key Takeaways
// Unbuffered channels block until a sender and receiver are both ready.
// The send goroutine gets blocked at ch <- "message" until receive reads the value.
// Once the receive goroutine reads from ch, the send goroutine unblocks and continues execution.
// Channel length (len(ch)) remains 0 because it's unbuffered—values don't stay in it.
