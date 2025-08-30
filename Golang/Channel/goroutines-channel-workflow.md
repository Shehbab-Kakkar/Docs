This Go program demonstrates the use of **goroutines** and **channels** for concurrent execution and communication between functions.

---

## 🧠 Program Explanation

### ✅ Code

```go
package main

import "fmt"

func sum(result chan int, num1 int, num2 int) {
	numResult := num1 + num2
	result <- numResult // Send result to the channel
}

func main() {
	result := make(chan int) // Create a channel of type int
	go sum(result, 4, 5)     // Start a goroutine to run sum()
	res := <-result          // Receive value from the channel
	fmt.Println(res)         // Print result
}
```

---

## 🔁 Step-by-Step Execution Flow

### 1. **Channel Creation**

```go
result := make(chan int)
```

* A channel of type `int` is created.
* Channels are used to **communicate between goroutines**.
* Think of it like a pipe: one goroutine sends a value into the pipe, and another receives it.

---

### 2. **Goroutine Started**

```go
go sum(result, 4, 5)
```

* The `sum` function is launched as a **goroutine** (a lightweight thread).
* It runs independently of `main()`.

---

### 3. **Inside Goroutine**

```go
numResult := num1 + num2
result <- numResult
```

* `num1` and `num2` are added (4 + 5 = 9).
* The result (`9`) is **sent into the channel** using `result <- numResult`.

---

### 4. **Receive Value in Main**

```go
res := <-result
```

* The main goroutine **waits to receive a value** from the channel.
* It **blocks until** it gets something. When the `sum()` function sends `9`, `main()` continues.

---

### 5. **Print the Result**

```go
fmt.Println(res)
```

* Output: `9`

---

## 📦 Channel Summary

| Operation | Symbol           | Meaning                                        |
| --------- | ---------------- | ---------------------------------------------- |
| Send      | `channel <- x`   | Send value `x` into `channel`                  |
| Receive   | `x := <-channel` | Receive value from `channel` and assign to `x` |

* Channels **synchronize** data between goroutines.
* They **block** the sender or receiver until the other side is ready.

---

## 🔄 Data Flow and Values

### Data Flow Diagram (Text Representation)

```text
Main Goroutine:               sum() Goroutine:
---------------               -----------------
make(chan int) ───┐
                  │
  go sum(result, 4, 5) ───────► [ numResult = 9 ]
                                  result <- 9
                  │
       res = <-result ◄─────────
                  │
          fmt.Println(res) ⇒ 9
```

---

## ✅ Summary of Concepts

| Concept         | Description                                                       |
| --------------- | ----------------------------------------------------------------- |
| **Goroutine**   | A lightweight thread to run functions concurrently                |
| **Channel**     | A way to send/receive values between goroutines                   |
| **Blocking**    | Channel operations block until both sender and receiver are ready |
| **Concurrency** | Both `main` and `sum` run at the same time via goroutines         |

---

Let me know if you want a **visual diagram** or if you'd like this expanded with buffered channels, select statements, or real-world use cases.
