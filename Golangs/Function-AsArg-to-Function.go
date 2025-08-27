package main

import (
	"fmt"
)

// Define a function type that takes two ints and returns an int
type opFunc func(int, int) int

// This function takes another function as an argument and returns an int
func calculate(f opFunc, a int, b int) int {
	return f(a, b)
}

// Example functions matching opFunc type
func add(x, y int) int {
	return x + y
}

func multiply(x, y int) int {
	return x * y
}

func main() {
	result1 := calculate(add, 3, 4)        // Pass 'add' function
	result2 := calculate(multiply, 5, 6)   // Pass 'multiply' function

	fmt.Println("Addition:", result1)       // Output: Addition: 7
	fmt.Println("Multiplication:", result2) // Output: Multiplication: 30
}
/*
Input for the program ( Optional )
STDIN
Output:

Addition: 7
Multiplication: 30
*/
