 **Go program with detailed inline comments**, including **explanations** and **data flow** between goroutines and channels:

---

```go
package main

import "fmt"

func main() {
	// Create an unbuffered int channel
	chan1 := make(chan int)

	// Create an unbuffered string channel
	chan2 := make(chan string)

	// Launch a goroutine to send data into chan1
	go func() {
		// Send the integer 10 into chan1
		// This send will block until main goroutine receives from chan1
		chan1 <- 10
	}()

	// Launch another goroutine to send data into chan2
	go func() {
		// Send the string "pong" into chan2
		// This send will block until main goroutine receives from chan2
		chan2 <- "pong"
	}()

	// Use a loop to receive from both channels exactly once
	for i := 0; i < 2; i++ {
		// select lets us wait on multiple channel operations
		select {
		// If chan1 has a value ready, receive it and bind to chan1Val
		case chan1Val := <-chan1:
			fmt.Println("received data from chan1:", chan1Val)

		// If chan2 has a value ready, receive it and bind to chan2Val
		case chan2Val := <-chan2:
			fmt.Println("received data from chan2:", chan2Val)
		}
	}
}
```

---

### 🔄 **Explanation of Data Flow**

1. **Main goroutine (starts first):**

   * Creates two unbuffered channels: `chan1` (int), `chan2` (string).
   * Starts two goroutines to send data into these channels.

2. **First goroutine**:

   * Tries to send `10` into `chan1`.
   * It **blocks** until the main goroutine receives from `chan1`.

3. **Second goroutine**:

   * Tries to send `"pong"` into `chan2`.
   * It **blocks** until the main goroutine receives from `chan2`.

4. **Back in the main goroutine**:

   * The `for` loop runs twice (`i := 0; i < 2`).
   * On each iteration, a `select` statement waits for **either** `chan1` or `chan2` to become ready.
   * Whichever channel is ready first (depending on goroutine scheduling), it receives the value and prints it.
   * Both channels are read exactly once over the two iterations.

---

### 🧠 Key Concepts

| Concept                | Description                                                                                 |
| ---------------------- | ------------------------------------------------------------------------------------------- |
| **Unbuffered Channel** | Send blocks until another goroutine is ready to receive.                                    |
| **Goroutine**          | Lightweight thread managed by Go runtime. Used to simulate concurrency.                     |
| **select**             | Waits for one of several channel operations to proceed. Only one `case` runs per iteration. |
| **Blocking**           | Both sending and receiving on unbuffered channels block until the opposite side is ready.   |

---

### 📌 Sample Output

Order may vary depending on which goroutine sends first:

```
received data from chan1: 10
received data from chan2: pong
```

or

```
received data from chan2: pong
received data from chan1: 10
```

