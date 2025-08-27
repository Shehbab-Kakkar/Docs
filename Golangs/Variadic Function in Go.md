
## What is a Variadic Function in Go?

A **variadic function** is a function that can accept zero, one, or multiple arguments of a specific type. With variadic functions, you don’t need to know in advance how many values will be passed to the function. This makes them very flexible for cases like summing numbers, joining strings, etc.

In Go, you define a variadic function by adding `...` before the type of the last parameter.

---

## Your Example Code

```go
package main
import "fmt"

// Variadic function: accepts any number of int arguments
func sum(num ...int) int {
    total := 0
    for _, n := range(num){
       total = total + n
    } 
    return total   
}

func main() {
   result := sum(34, 55, 67, 78) // Passing four integers
   fmt.Println(result)           // Output: 234
}
```

---

## Explanation Step-by-Step

### 1. Function Definition

```go
func sum(num ...int) int {
    // function body
}
```
- `sum` is a variadic function.
- `num ...int` means you can pass any number of `int` values to this function.
- Inside the function, `num` is a **slice** of `int` (`[]int`). Even if you pass individual numbers, Go automatically packages them into a slice.

### 2. Using the Variadic Parameter

```go
total := 0
for _, n := range(num){
   total = total + n
}
```
- You loop over the `num` slice using `range`.
- Each `n` is one of the passed-in integers.
- You add each `n` to `total`.

### 3. Returning the Result

```go
return total
```
- After the loop, `total` contains the sum of all the arguments passed to the function.

### 4. Calling the Variadic Function

```go
result := sum(34, 55, 67, 78)
```
- You can pass as many integers as you want.
- The function adds them up and returns the result.

---

## Key Points About Variadic Functions

1. **Zero or More Arguments:** You can call `sum()` with no values, one value, or many values.
2. **Slice as Argument:** Internally, the variadic parameter is a slice (`[]int` in this case).
3. **Passing a Slice:** If you already have a slice, you can pass it using `...`:
   ```go
   nums := []int{10, 20, 30}
   result := sum(nums...) // Spread the slice into individual arguments
   ```
4. **Only the Last Parameter Can Be Variadic:** You can have other parameters, but only the last one can be variadic.

---

## Example with No Arguments

```go
fmt.Println(sum()) // Output: 0 (since no numbers are passed)
```

---

## Summary

- Variadic functions are perfect when you want to let users pass any number of arguments.
- The `...` syntax before the type lets Go know to collect all extra arguments into a slice.
- They make your code flexible and concise.

**Your example demonstrates how easy it is to sum any number of integers using a variadic function in Go!**
