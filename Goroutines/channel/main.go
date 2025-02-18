package main

import "fmt"

func main() {
	// Create an unbuffered channel of type int
	ch := make(chan int)

	// Define an anonymous function to multiply two numbers and send the result through the channel
	mult := func(x, y int) {
		res := x * y // Multiply x and y
		fmt.Println("Inside goroutine, sending value to channel:", res)
		ch <- res // Send the result to the channel
	}

	// Start a goroutine to execute the multiplication function
	go mult(10, 10)

	// Receive the value from the channel
	val, ok := <-ch

	// Print the type and value of the channel
	fmt.Printf("Type of value of ch is: %T\n", ch)
	fmt.Printf("The value of ch is: %v\n", ch)

	// Print the received value and the status of the channel
	fmt.Printf("The resulting value from the goroutine is: %v\n", val)
	fmt.Printf("The value of ok is: %v\n", ok)
}

// output
// Inside goroutine, sending value to channel: 100
// Type of value of ch is: chan int
// The value of ch is: 0xc0000120c0  // (Memory address of the channel)
// The resulting value from the goroutine is: 100
// The value of ok is: true (true indicate value is recieved by write operation, false if value is recieved after channel is closed.)

// The main function does not proceed past <-ch (channel receive) until it gets a value from the channel.
// Channels in Go block by default when sending and receiving data.
// This means the main goroutine pauses execution at val, ok := <-ch until the value is available.

// Step-by-Step Execution:
// Create the channel (ch := make(chan int))

// An unbuffered channel is created.
// Launch the Goroutine (go mult(10, 10))

// The multiplication function runs in the background asynchronously.
// Main Goroutine Reaches val, ok := <-ch

// At this point, the main goroutine is blocked because it is waiting to receive a value from ch.
// It cannot move forward until the goroutine sends a value.
// The Goroutine Executes and Sends 100 (ch <- res)

// Inside mult(10, 10), res = 10 * 10 = 100.
// The result (100) is sent to ch, unblocking the main goroutine.
// Main Goroutine Resumes Execution

// Now that a value is available in ch, val gets assigned 100, and execution continues.

// Key Concept: Channels Block Until Communication Happens
// Sending (ch <- res) blocks until there is a receiver (<-ch).
// Receiving (<-ch) blocks until a value is sent.
