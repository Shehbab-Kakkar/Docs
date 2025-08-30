```go
package main

import (
    "fmt"
)

func main() {
    ch := make(chan int)

    go func() {
        ch <- 42 // blocks until main goroutine receives
    }()

    x := <-ch // receives 42, unblocks the sender
    fmt.Println(x) // Output: 42
}
```

Question: If we do go ch <- 42 will it works?

Great question!

You're asking if this line:

```go
go ch <- 42
```

...will work the same as this:

```go
go func() {
    ch <- 42
}()
```

### Short Answer:

**No**, `go ch <- 42` **will not compile**.

### Why?

Because Go expects the argument to a `go` statement to be a **function call**. `ch <- 42` is a **send statement**, not a function call — and Go won't let you launch a goroutine from it directly.

### Compile Error Example:

If you try:

```go
go ch <- 42
```

You'll get a compile-time error like:

```
syntax error: unexpected <-, expecting (
```

### Correct Approach:

You need to **wrap** the send operation inside a function (anonymous or named), like this:

```go
go func() {
    ch <- 42
}()
```

### Summary:

* ✅ `go func() { ch <- 42 }()` → works
* ❌ `go ch <- 42` → compile-time error

Let me know if you want to see what happens with buffered channels or how the runtime handles it!
