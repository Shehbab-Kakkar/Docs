package main

import (
	"fmt"
	"time"
)

// order struct
type order struct {
	id       string
	amount   float32
	status   string
	createAt time.Time // nanosecond precision
}

// Receiver method to update the status
func (o *order) updateStatus(newStatus string) {
	o.status = newStatus
}

func main() {
	// Step 1: Declare and initialize a new order
	var o = order{
		id:     "1",
		amount: 50.00,
		status: "received",
	}

	// Step 2: Assign current time to the createAt field
	o.createAt = time.Now()

	// Step 3: Print order before status update
	fmt.Println("Before update:", o)

	// Step 4: Call method to update the status
	o.updateStatus("processing")

	// Step 5: Print order after status update
	fmt.Println("After update:", o)
}
