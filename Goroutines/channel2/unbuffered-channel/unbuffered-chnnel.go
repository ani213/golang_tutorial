package main

import (
	"fmt"
	"time"
)

func sendEmail(email chan string, done chan bool) {
	defer func() {
		done <- true // Signal that the email sending is done
	}()
	// Simulate sending emails
	for email := range email {
		fmt.Println("sending eamil to:-", email) // Simulate sending an email
	}
}

func main() {
	emailChannel := make(chan string, 3)
	done := make(chan bool)

	go sendEmail(emailChannel, done) // Start the goroutine to send emails

	for i := 0; i <= 50; i++ {
		emailChannel <- fmt.Sprintf("%d@example.com", i)
		time.Sleep(time.Second)
	}
	close(emailChannel) // Close the channel to indicate no more emails will be sent
	fmt.Println("All emails sent, waiting for goroutine to finish...")
	<-done // Wait for the goroutine to finish sending emails

	// emailChannel <- "1@example.com"
	// emailChannel <- "2@example.com"
	// fmt.Println(<-emailChannel)         // Receive and print the first email
	// fmt.Println(<-emailChannel)         // Receive and print the second email
	// fmt.Println("Channel is empty now") // Indicate that the channel is empty
	// close(emailChannel)                 // Close the channel to indicate no more values will be sent
	// fmt.Println("Channel closed")       // Print a message indicating the channel is closed
}
