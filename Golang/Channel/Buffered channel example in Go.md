 Go program with **detailed comments explaining each part of the code** and **how data flows between goroutines using channels**:

```go
// Buffered channel example in Go
package main

import (
	"fmt"
	"time"
)

// emailSender simulates sending emails from a buffered channel.
// It receives emails from the channel and prints "sending email to ...".
// Once the channel is closed and all emails are sent, it signals completion through the 'done' channel.
func emailSender(emailChan chan string, done chan bool) {
	defer func() { done <- true }() // Signal main goroutine when done
	for email := range emailChan {  // Receives from emailChan until it's closed
		fmt.Println("sending email to", email)
		time.Sleep(time.Second) // Simulate delay in sending
	}
}

func main() {
	// Create a buffered channel of strings with capacity 5
	emailChan := make(chan string, 5)

	// Create an unbuffered channel to signal when emailSender is done
	done := make(chan bool)

	// Start emailSender goroutine
	go emailSender(emailChan, done)

	// Send 5 email addresses into the buffered channel
	for i := 0; i < 5; i++ {
		emailChan <- fmt.Sprintf("%d@gmail.com", i) // Send email to emailChan
	}

	fmt.Println("done sending") // Message after all emails have been queued into the channel

	// Close the emailChan to indicate that no more emails will be sent
	close(emailChan)

	// Wait for the emailSender to finish processing all emails
	<-done
}
```

---

### 🔄 **Data Flow and Execution Order**

1. **Main goroutine (main function)** starts and:

   * Creates two channels:

     * `emailChan`: Buffered channel with size 5. Used to pass emails to be sent.
     * `done`: Unbuffered channel used for synchronization — waits for the `emailSender` to finish.

2. **Goroutine starts**:

   * `emailSender(emailChan, done)` is launched in a new goroutine and starts waiting for data from `emailChan`.

3. **Main function sends 5 email addresses** into `emailChan`:

   * Since the channel is **buffered with size 5**, the sends are **non-blocking** until the buffer is full.
   * All 5 email addresses are sent quickly and stored in the buffer.

4. `fmt.Println("done sending")` prints, even before emails are sent, because the sending to the channel is fast (buffered).

5. `emailChan` is **closed** to indicate to `emailSender` that no more values will come.

6. **emailSender goroutine reads** from `emailChan` one by one:

   * For each email, it prints "sending email to ..." and waits 1 second (simulating a delay).
   * It exits when the channel is closed and empty.

7. After all emails are processed, `emailSender` sends `true` to the `done` channel.

8. **Main goroutine blocks on `<-done`** until it receives the signal from `emailSender`, then exits.

---

### 📌 Summary of Channel Use

| Channel     | Type       | Purpose                                   |
| ----------- | ---------- | ----------------------------------------- |
| `emailChan` | Buffered   | Pass email strings from main to goroutine |
| `done`      | Unbuffered | Notify main when goroutine is finished    |

---

