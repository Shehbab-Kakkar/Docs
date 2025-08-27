package main

import (
	"fmt"
	"time"
)

// order struct definition
type order struct {
	id       string
	amount   float32
	status   string
	createAt time.Time
}

// Constructor function that returns a *pointer* to order
func NewOrder(id string, amount float32, status string) *order {
	return &order{
		id:       id,
		amount:   amount,
		status:   status,
		createAt: time.Now(), // set creation time at instantiation
	}
}

// Method with pointer receiver to update the order status
func (o *order) updateStatus(newStatus string) {
	o.status = newStatus
}

func main() {
	// Step 1: Create a new order using the constructor
	o := NewOrder("1", 50.00, "received")

	// Step 2: Print order before status update
	fmt.Println("Before update:", *o)

	// Step 3: Update the status
	o.updateStatus("processing")

	// Step 4: Print order after status update
	fmt.Println("After update:", *o)
}
