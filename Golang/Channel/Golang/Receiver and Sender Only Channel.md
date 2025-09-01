Go program** followed by a **step-by-step explanation** of how the sender and receiver are structured, how data flows through the program, and what makes this format clean and efficient.

---

## ✅ Updated Go Program with Receiver and Sender Separated

```go
package main

import (
	"fmt"
	"time"
)

// emailReceiver receives emails from the emailChan and simulates sending them.
// It runs as a separate goroutine and exits when the channel is closed.
func emailReceiver(emailChan <-chan string, done chan<- bool) {
	defer func() { done <- true }() // Notify main() when done
	for email := range emailChan {
		fmt.Println("sending email to", email)
		time.Sleep(time.Second) // Simulate delay in sending
	}
}

// emailSender sends a specified number of email addresses into the emailChan.
// It runs as a separate goroutine and closes the channel after all emails are sent.
func emailSender(emailChan chan<- string, count int) {
	for i := 0; i < count; i++ {
		email := fmt.Sprintf("%d@gmail.com", i)
		fmt.Println("queueing email:", email)
		emailChan <- email // Send email into channel
	}
	close(emailChan) // Signal that no more emails will be sent
}

func main() {
	emailChan := make(chan string, 5) // Buffered channel for email strings
	done := make(chan bool)           // To notify when all emails are processed

	// Start the receiver (consumer) goroutine
	go emailReceiver(emailChan, done)

	// Start the sender (producer) goroutine
	go emailSender(emailChan, 5)

	// Wait until receiver signals that all emails have been processed
	<-done
	fmt.Println("All emails processed.")
}
```

---

## 🧠 Explanation: Sender and Receiver Format

### 🔹 **Sender Format (`emailSender`)**

```go
func emailSender(emailChan chan<- string, count int)
```

* `chan<- string`: This is a **send-only channel**. The function can only **send** data into `emailChan`.
* It loops `count` times and generates emails like `"0@gmail.com"`, `"1@gmail.com"`, etc.
* Each email is sent into `emailChan` using:

  ```go
  emailChan <- email
  ```
* After all emails are sent, the channel is **closed** using:

  ```go
  close(emailChan)
  ```

### 🔹 **Receiver Format (`emailReceiver`)**

```go
func emailReceiver(emailChan <-chan string, done chan<- bool)
```

* `<-chan string`: This is a **receive-only channel**. The function can only **read** from `emailChan`.
* `chan<- bool`: Used to **send a signal back** to `main()` once processing is done.
* It reads from the channel using:

  ```go
  for email := range emailChan
  ```

  This loop runs until the channel is **closed** and all emails are read.
* For each email, it prints a "sending" message and simulates a delay.
* Once done, it sends `true` to the `done` channel to notify `main()`.

---

## 🔄 Data Flow in the Program

Here’s a step-by-step flow of how data moves:

```
                +-----------------+
                |    main()       |
                +-----------------+
                        |
              creates channels:
              emailChan, done
                        |
             starts two goroutines:
              /                    \
     +----------------+      +------------------+
     |  emailSender   |      |  emailReceiver   |
     +----------------+      +------------------+
            |                        |
   queueing email into        receiving email from
       emailChan                  emailChan
       (buffered)               (blocks if empty)
            |                        |
    close(emailChan)        loop exits when closed
            |                        |
                               sends true to done
                        (signals that it's finished)
                        |
                 main() waits for <-done
                        |
             prints "All emails processed."
```

---

## 🧪 Sample Output (order may vary)

```
queueing email: 0@gmail.com
queueing email: 1@gmail.com
queueing email: 2@gmail.com
queueing email: 3@gmail.com
queueing email: 4@gmail.com
sending email to 0@gmail.com
sending email to 1@gmail.com
sending email to 2@gmail.com
sending email to 3@gmail.com
sending email to 4@gmail.com
All emails processed.
```

---

## ✅ Advantages of This Format

| Benefit                   | Why it Matters                                    |
| ------------------------- | ------------------------------------------------- |
| Clear separation of roles | Easier to debug, extend, or reuse sender/receiver |
| Directional channels      | Enforces correct usage of channels                |
| Graceful termination      | Receiver stops when channel is closed             |
| Buffered channel          | Sender can send a few items without blocking      |
| Done signaling            | Clean way to know when the job is finished        |

---

Let me know if you'd like to:

* Add multiple receivers or senders
* Handle failures or timeouts
* Turn this into a real-world email queue system

Happy to extend it further!
