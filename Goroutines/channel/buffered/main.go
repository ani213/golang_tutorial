package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	// Create a buffered channel with a capacity of 3.
	// This means it can store up to 3 messages before blocking further sends.
	ch := make(chan string, 3)

	// Start the 'send' goroutine to send messages into the channel.
	go send(ch)

	// Prompt the user to press enter to continue, ensuring that sends complete first.
	fmt.Println("\nPress Enter key to continue...")
	fmt.Scanln()

	// Start the 'receive' goroutine to receive messages from the channel.
	go receive(ch)

	// Sleep for 1 second to allow the goroutines to complete before the program exits.
	time.Sleep(time.Second * 1)
}

// send sends messages into the channel.
func send(ch chan string) {
	// Print the initial length of the channel (should be 0).
	fmt.Printf("Channel length before sending anything: %v\n", len(ch))

	// Loop to send 3 messages into the channel.
	for i := 1; i < 4; i++ {
		// Create a message string dynamically.
		msg := "message" + strconv.Itoa(i)

		// Print the message being sent.
		fmt.Printf("Sending: %v\n", msg)

		// Send the message into the channel.
		// Since the channel has a buffer size of 3, these will not block.
		ch <- msg

		// Print the updated length of the channel after sending.
		fmt.Printf("Channel length: %v\n", len(ch))
	}
}

// receive receives messages from the channel.
func receive(ch chan string) {
	// Loop to receive 3 messages.
	for i := 0; i < 3; i++ {
		// Receive a message from the channel and print it.
		fmt.Printf("Received: %v\n", <-ch)

		// Print the updated length of the channel after receiving.
		fmt.Printf("Channel length: %v\n", len(ch))
	}
}

// output -
// Press Enter key to continue...
// Channel length before sending anything: 0
// Sending: message1
// Channel length: 1
// Sending: message2
// Channel length: 2
// Sending: message3
// Channel length: 3

// Received: message1
// Channel length: 2
// Received: message2
// Channel length: 1
// Received: message3
// Channel length: 0
