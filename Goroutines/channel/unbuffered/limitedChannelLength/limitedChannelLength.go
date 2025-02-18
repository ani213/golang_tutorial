package limitedChannelLength

import (
	"fmt"
	"time"
)

// LimitedChannelLength demonstrates the behavior of a buffered channel with limited capacity.
func LimitedChannelLength() {
	// Create a buffered channel with a capacity of 1.
	// This means it can hold only one message at a time before blocking further sends.
	ch := make(chan string, 1)

	// Start the 'send1' goroutine, which will attempt to send messages into the channel.
	go send1(ch)

	// Print the length of the channel after attempting to send.
	// Initially, it should be 0 because the 'send1' goroutine hasn't executed yet.
	fmt.Printf("channel length: %v\n", len(ch))

	// Start the 'receive1' goroutine to receive messages after a delay.
	go receive1(ch)

	// Sleep to allow the goroutines to complete execution before the function exits.
	// This prevents premature termination of the program.
	time.Sleep(time.Second * 2)
}

// send1 sends messages into the channel.
func send1(ch chan string) {
	fmt.Println("blocking send goroutine...")

	// Sending the first message into the buffered channel.
	// Since the channel has a capacity of 1, this message will fit in the buffer immediately.
	ch <- "message1"

	// Attempting to send another message.
	// Since the channel is already full (capacity is 1), this will block until a receiver reads the first message.
	ch <- "message2"

	// This line executes only after the second message has been successfully sent.
	fmt.Println("send goroutine unblocked...")
}

// receive1 receives messages from the channel.
func receive1(ch chan string) {
	// Simulate some delay before receiving the message.
	time.Sleep(time.Second * 1)

	// Print the channel length before consuming any messages.
	// It should be 1 because only one message fits in the buffer at a time.
	fmt.Printf("channel length: %v\n", len(ch))

	// Receive the first message from the channel, unblocking the 'send1' goroutine.
	fmt.Printf("received: %v\n", <-ch)

	// Receive the second message from the channel.
	// The second message was previously blocked, but now it is available.
	fmt.Printf("received: %v\n", <-ch)
}

// output -
// channel length: 0
// blocking send goroutine...
// channel length: 1
// received: message1
// received: message2
// send goroutine unblocked...

// or
// channel length: 0
// blocking send goroutine...
// channel length: 1
// send goroutine unblocked...
// received: message1
// received: message2

// Key Takeaways
// Buffered channels allow non-blocking sends up to their capacity.
// Sending beyond the capacity blocks the sender until a message is received.
// Receiving messages unblocks blocked send operations.
// Execution timing matters—goroutines execute asynchronously, so delays help demonstrate behavior.
