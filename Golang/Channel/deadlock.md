Let's analyze your Go program and the resulting **deadlock error**:

```go
package main

import "fmt"

func main() {
	messageChan := make(chan string)
	messageChan <- "Message from channel"  // (1)
	msg := <-messageChan                   // (2)
	fmt.Println(msg)
}
```

### What happens, step by step?

1. **`messageChan := make(chan string)`**
   - Creates an **unbuffered** channel of type `string`.

2. **`messageChan <- "Message from channel"`**
   - The main goroutine tries to **send** a string into the channel.
   - **Unbuffered channels** require that both a sender and receiver are ready at the same time. If not, the sending or receiving goroutine blocks (waits).

3. **`msg := <-messageChan`**
   - The main goroutine will only reach this line **after** the send completes. But the send will not complete until something is ready to receive.
   - **But** since both lines are in the same (main) goroutine, the send blocks **forever**—waiting for a receiver that never comes.

4. As a result, **the program deadlocks**:  
   - The send operation at line (1) blocks the only goroutine (main), so it never reaches the receive operation at line (2).
   - Go detects that all goroutines are asleep (blocked) and the program cannot proceed.

### Why this deadlock occurs

- **Unbuffered channels** need the sender and receiver to be ready at the same time.  
- Both operations can only proceed if one is in a separate goroutine:

#### Example 1: Using a goroutine for sending
```go
go func() {
    messageChan <- "Message from channel"
}()
msg := <-messageChan
fmt.Println(msg)
```

#### Example 2: Using a buffered channel
```go
messageChan := make(chan string, 1) // Buffer size 1
messageChan <- "Message from channel"
msg := <-messageChan
fmt.Println(msg)
```

### Summary Table

| Channel Type         | Sender/Receiver in Same Goroutine? | Deadlock? |
|----------------------|:----------------------------------:|:---------:|
| Unbuffered           | Yes                                | Yes       |
| Unbuffered           | No (separate goroutine)            | No        |
| Buffered (capacity 1)| Yes                                | No        |

---

### **Conclusion**

- The error occurs because you are trying to send and receive on an unbuffered channel in the same goroutine, which leads to a deadlock.
- Use a separate goroutine for sending/receiving, or use a **buffered** channel to fix it.

Let me know if you want to see the corrected code!
