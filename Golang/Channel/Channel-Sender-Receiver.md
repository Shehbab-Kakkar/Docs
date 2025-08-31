Great! Let's break down your program and clearly explain the **data flow**, the **sender** and **receiver** roles, and how **blocking** behavior works in this context.

---

## ✅ PROGRAM OVERVIEW

### 🔹 Code

```go
package main

import (
	"fmt"
)

// Function to add two numbers and send the result into a channel
func sum(result chan int, num1 int, num2 int) {
	numResult := num1 + num2
	result <- numResult // Send result into the channel (sender)
}

func main() {
	result := make(chan int)      // Create an unbuffered channel
	go sum(result, 4, 5)          // Launch sum() as a goroutine (concurrent execution)
	res := <-result               // Receive value from the channel (receiver) - blocks until available
	fmt.Println(res)              // Print the received result
}
```

---

## 🔄 DATA FLOW EXPLANATION

### Step-by-step Execution

1. **`main()` creates a channel:**

   ```go
   result := make(chan int)
   ```

   * This is an **unbuffered channel** for sending/receiving `int` values.
   * No data can be sent through this channel unless there’s a receiver ready.

2. **Start goroutine to compute sum:**

   ```go
   go sum(result, 4, 5)
   ```

   * A **goroutine** is launched to run the `sum` function concurrently.
   * Inside `sum`, the result of `4 + 5` (which is `9`) is sent to the channel:

     ```go
     result <- numResult
     ```

3. **Receive the result in `main()`:**

   ```go
   res := <-result
   ```

   * This line **blocks** until the value is received from the `sum` goroutine.
   * Once the goroutine sends `9` into the channel, this line completes, and `res` is assigned the value `9`.

4. **Print the result:**

   ```go
   fmt.Println(res) // Output: 9
   ```

---

## 🎯 ROLES: SENDER vs RECEIVER

| Role         | Function            | Channel Operation     |
| ------------ | ------------------- | --------------------- |
| **Sender**   | `sum()` (goroutine) | `result <- numResult` |
| **Receiver** | `main()`            | `res := <-result`     |

* The **sender** is the `sum()` function, which **sends** data into the channel.
* The **receiver** is the `main()` function, which **receives** data from the channel.

---

## ⚠️ BLOCKING BEHAVIOR

* Since the channel is **unbuffered**, sending and receiving must happen at the same time:

  * The goroutine in `sum()` **blocks** on:

    ```go
    result <- numResult
    ```

    until `main()` is ready to receive.
  * `main()` **blocks** on:

    ```go
    res := <-result
    ```

    until the `sum()` function sends the value.

🔁 These two operations **synchronize** through the channel.

---

## 📊 FLOW DIAGRAM

```plaintext
main()                                sum() goroutine
  |                                      |
  |                                     compute 4 + 5 = 9
  |                                      |
  |<----------- receive 9  <-------------|  result <- 9
  |                                      |
fmt.Println(9)                           |
```

* Data flows **from `sum()` to `main()`** via the `result` channel.

---

## ✅ OUTPUT

```
9
```

---

Let me know if you'd like this example modified to use buffered channels, multiple goroutines, or error handling.
