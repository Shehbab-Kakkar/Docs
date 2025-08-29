```go
package main

import (
	"fmt"
	"sync"
)

func worker(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("worker %d started\n", i)
	fmt.Printf("worker %d end\n", i)

}
func main() {
	//fmt.Println("Explore goroutine started")
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1) //increment the WaitGroup Counter
		go worker(i, &wg)
	}

	//Wait for all workers to finish
	wg.Wait() //This responsible for all go routines
	fmt.Println("Workers task complete")
}

```
 This explanation covers the concurrency model, use of goroutines and `sync.WaitGroup`, and a line-by-line overview for clarity.

---

## 🧵 Go Concurrency Example with Goroutines and WaitGroup

This Go program demonstrates **concurrent execution** using **goroutines** and the **`sync.WaitGroup`** to coordinate multiple worker tasks.

### 🔍 Overview

The main goal of this program is to:

* Launch multiple **worker functions** concurrently using goroutines.
* Use a **`WaitGroup`** to wait for all goroutines to complete before continuing the main function.

---

### 📄 Code

```go
package main

import (
	"fmt"
	"sync"
)

func worker(i int, wg *sync.WaitGroup) {
	defer wg.Done() // Decrements the counter when the goroutine completes
	fmt.Printf("worker %d started\n", i)
	fmt.Printf("worker %d end\n", i)
}

func main() {
	var wg sync.WaitGroup // Initialize a WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1) // Increments the WaitGroup counter by 1 for each worker
		go worker(i, &wg) // Launches a worker goroutine
	}

	wg.Wait() // Blocks until the counter goes back to 0 (all workers are done)
	fmt.Println("Workers task complete")
}
```

---

### 🧠 How It Works

#### ✅ `sync.WaitGroup`

* A **WaitGroup** is used to wait for a collection of goroutines to finish executing.
* It has three main methods:

  * `Add(n int)` – Increments the counter by `n`
  * `Done()` – Decrements the counter by 1 (usually called using `defer`)
  * `Wait()` – Blocks until the counter becomes 0

#### ⚙️ Execution Flow

1. **WaitGroup Initialization**:

   * A `WaitGroup` named `wg` is created to track when all workers are finished.

2. **Spawning Goroutines**:

   * A loop runs from 1 to 3. For each iteration:

     * `wg.Add(1)` increases the counter.
     * `go worker(i, &wg)` launches a new goroutine that runs the `worker` function concurrently.

3. **Worker Function**:

   * Each worker prints a start message.
   * Then it prints an end message.
   * Finally, it calls `wg.Done()` via `defer` to signal its completion.

4. **Wait for Completion**:

   * `wg.Wait()` ensures that the `main()` function does not exit until all 3 workers have called `Done()`.

5. **Final Output**:

   * Once all goroutines finish, the main function prints `"Workers task complete"`.

---

### 🧪 Sample Output

```bash
go run goroute/main.go 
worker 3 started
worker 3 end
worker 1 started
worker 1 end
worker 2 started
worker 2 end
Workers task complete
```

> ⚠️ **Note**: The order of worker output may vary on each run. Goroutines are scheduled independently and can execute in any order.

---

### 📌 Key Takeaways

* Goroutines allow functions to run concurrently without blocking the main thread.
* `sync.WaitGroup` is essential to coordinate and wait for multiple goroutines.
* Always use `defer wg.Done()` at the beginning of the goroutine to ensure proper decrementing even if the function panics or returns early.
* Go’s concurrency model is lightweight and powerful for building scalable, performant applications.

---

Let me know if you'd like a visual diagram or want to extend this with error handling or context cancellation!
