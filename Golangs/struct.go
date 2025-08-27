package main

import (
	"fmt"
	"time"
)

// orde struct
type order struct {
	id       string
	amount   float32
	status   string
	createAt time.Time //nanosecond precision
}

func main() {
	var order = order{
		id:     "1",
		amount: 50.00,
		status: "received",
	}
	order.createAt = time.Now()
	fmt.Println(order)
	//fmt.Println("Order created at:", order.createAt.Format("2006-01-02 15:04:05"))

}
/*
output:
{1 50 received {13989286089280496184 39332 0x5709e0}}
Order created at: 2025-08-27 17:22:25
*/
