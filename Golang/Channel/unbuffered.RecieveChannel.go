package main

import "fmt"

func main() {
	// Create an unbuffered int channel
	chan1 := make(chan int)

	// Create an unbuffered string channel
	chan2 := make(chan string)

	// Launch a goroutine to send data into chan1
	go func() {
		// Send the integer 10 into chan1
		// This send will block until main goroutine receives from chan1
		chan1 <- 10
	}()

	// Launch another goroutine to send data into chan2
	go func() {
		// Send the string "pong" into chan2
		// This send will block until main goroutine receives from chan2
		chan2 <- "pong"
	}()

	// Use a loop to receive from both channels exactly once
	for i := 0; i < 2; i++ {
		// select lets us wait on multiple channel operations
		select {
		// If chan1 has a value ready, receive it and bind to chan1Val
		case chan1Val := <-chan1:
			fmt.Println("received data from chan1:", chan1Val)

		// If chan2 has a value ready, receive it and bind to chan2Val
		case chan2Val := <-chan2:
			fmt.Println("received data from chan2:", chan2Val)
		}
	}
}
