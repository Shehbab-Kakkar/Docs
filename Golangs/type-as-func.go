package main

import (
    "fmt"
)

// Declare a function type
type opFunc func(int, int) int

// Define a function matching that type
func add(x, y int) int {
    return x + y
}

func main() {
    // Assign the function to a variable of that type
    var f opFunc = add

    // Call the function via the variable
    result := f(2, 3)

    // Print the result
    fmt.Println("Result:", result)
}
