package main

import (
    "fmt"
)

// Define a function type
type opFunc func(int, int) int

// Define operations
func add(x, y int) int {
    return x + y
}

func subtract(x, y int) int {
    return x - y
}

func multiply(x, y int) int {
    return x * y
}

func divide(x, y int) int {
    if y == 0 {
        fmt.Println("Error: Division by zero")
        return 0
    }
    return x / y
}

func remainder(x, y int) int {
    if y == 0 {
        fmt.Println("Error: Division by zero")
        return 0
    }
    return x % y
}

// Square: returns x squared (ignores y)
func square(x, _ int) int {
    return x * x
}

func main() {
    var f opFunc

    // Test data
    x := 10
    y := 3

    // Addition
    f = add
    fmt.Println("Add:", f(x, y))

    // Subtraction
    f = subtract
    fmt.Println("Subtract:", f(x, y))

    // Multiplication
    f = multiply
    fmt.Println("Multiply:", f(x, y))

    // Division
    f = divide
    fmt.Println("Divide:", f(x, y))

    // Remainder
    f = remainder
    fmt.Println("Remainder:", f(x, y))

    // Square
    f = square
    fmt.Println("Square of", x, "is:", f(x, 0)) // y is ignored in square
}
