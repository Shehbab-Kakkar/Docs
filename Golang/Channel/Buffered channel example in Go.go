// Buffered channel example in Go
package main

import (
	"fmt"
	"time"
)

// emailSender simulates sending emails from a buffered channel.
// It receives emails from the channel and prints "sending email to ...".
// Once the channel is closed and all emails are sent, it signals completion through the 'done' channel.
func emailSender(emailChan chan string, done chan bool) {
	defer func() { done <- true }() // Signal main goroutine when done
	for email := range emailChan {  // Receives from emailChan until it's closed
		fmt.Println("sending email to", email)
		time.Sleep(time.Second) // Simulate delay in sending
	}
}

func main() {
	// Create a buffered channel of strings with capacity 5
	emailChan := make(chan string, 5)

	// Create an unbuffered channel to signal when emailSender is done
	done := make(chan bool)

	// Start emailSender goroutine
	go emailSender(emailChan, done)

	// Send 5 email addresses into the buffered channel
	for i := 0; i < 5; i++ {
		emailChan <- fmt.Sprintf("%d@gmail.com", i) // Send email to emailChan
	}

	fmt.Println("done sending") // Message after all emails have been queued into the channel

	// Close the emailChan to indicate that no more emails will be sent
	close(emailChan)

	// Wait for the emailSender to finish processing all emails
	<-done
}
