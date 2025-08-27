package main

import (
    "fmt"
)

// Operation function type
type opFunc func(int, int) int

// Define math functions
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

// calculate takes another function as argument
func calculate(x, y int, op opFunc) int {
    return op(x, y)
}

func main() {
    x := 20
    y := 5

    fmt.Println("Add:", calculate(x, y, add))
    fmt.Println("Subtract:", calculate(x, y, subtract))
    fmt.Println("Multiply:", calculate(x, y, multiply))
    fmt.Println("Divide:", calculate(x, y, divide))
}
