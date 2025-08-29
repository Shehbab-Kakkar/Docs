package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("Say Hello Function")
	time.Sleep(10000 * time.Millisecond)
}
func sayHi() {
	fmt.Println("Say Hi 3")
	time.Sleep(10000 * time.Millisecond)
}
func main() {
	fmt.Println("Main function")
	go sayHello()
	go sayHi()
	time.Sleep(10000 * time.Millisecond)
}


/*
Main function
Say Hello Function
Say Hi 3
*/
