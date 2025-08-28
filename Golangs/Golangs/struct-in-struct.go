package main

import (
	"fmt"
	"time"
)

type customer struct {
	name  string
	phone string
}
type order struct {
	id       string
	amount   float32
	status   string
	createAt time.Time
	customer
}

func main() {
	newOrder := order{
		id:       "1",
		amount:   30,
		status:   "received",
		createAt: time.Now(),
		customer: customer{
			name:  "john",
			phone: "2344354646",
		},
	}

	fmt.Println(newOrder)
}
/*

*/
