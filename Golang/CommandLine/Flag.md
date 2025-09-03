In Go (Golang), the `flag` package is used for command-line flag parsing. It allows you to define flags (options) that can be passed to your program when it is executed from the command line. These flags are typically used for configuration—such as specifying file paths, enabling debug mode, or setting port numbers.

## How `flag` Works

1. **Define flags** using functions like `flag.String`, `flag.Int`, or `flag.Bool`.
2. **Parse flags** by calling `flag.Parse()`.
3. **Use the flag values** in your program.

## Example: Using `flag` Package in Go

Let's create a simple Go program that takes three command-line flags:

- `name` (string): The user's name
- `age` (int): The user's age
- `verbose` (bool): Enable verbose mode

```go
package main

import (
    "flag"
    "fmt"
)

func main() {
    // Define flags
    name := flag.String("name", "World", "your name")
    age := flag.Int("age", 0, "your age")
    verbose := flag.Bool("verbose", false, "enable verbose mode")

    // Parse the flags
    flag.Parse()

    // Use the flags
    if *verbose {
        fmt.Println("Verbose mode is enabled.")
    }
    fmt.Printf("Hello, %s! You are %d years old.\n", *name, *age)
}
```

### How to Run

Build and run the program from the command line:

```sh
go run main.go -name=Alice -age=30 -verbose
```

**Output:**
```
Verbose mode is enabled.
Hello, Alice! You are 30 years old.
```

### Explanation

- `flag.String("name", "World", "your name")` defines a string flag called `name` with a default value `"World"` and a description `"your name"`.
- `flag.Int("age", 0, "your age")` defines an integer flag called `age` with a default value `0`.
- `flag.Bool("verbose", false, "enable verbose mode")` defines a boolean flag.
- `flag.Parse()` processes the command-line arguments.
- The values are pointers, so you must dereference them (e.g., `*name`).

## Summary

The `flag` package is a convenient way to handle command-line options in Go programs. It helps you write flexible and configurable command-line tools.
