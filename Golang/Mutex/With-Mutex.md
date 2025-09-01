Great! You're now dealing with **concurrency** in Go — an important and powerful concept.

This program introduces `**sync.Mutex**` to handle **race conditions** when multiple goroutines try to modify shared data concurrently.

---

## ✅ What is a Mutex?

### 🔒 **Mutex = Mutual Exclusion**

A **Mutex** is a **lock** used to prevent multiple goroutines from **accessing/modifying the same data at the same time**.

In Go, a `sync.Mutex` is used to protect **critical sections** of code — parts that access shared variables — so that only **one goroutine at a time** can execute them.

---

## 🔥 The Problem Without a Mutex

In your example, 100 goroutines call `myPost.inc()` at the same time.

Without a `Mutex`, all these goroutines could try to read/modify `myPost.views` **simultaneously**, causing a **race condition**.

Race conditions result in:

* Unpredictable behavior
* Incorrect or inconsistent results
* Hard-to-reproduce bugs

For example, without the Mutex, the final value of `views` might be **less than 100**, even though `100` goroutines incremented it.

---

## ✅ What Mutex Solves

### Critical Section in Your Code:

```go
p.views += 1
```

This **read-modify-write** operation **must not be interrupted**, otherwise two goroutines may read the same value and both write back the same incremented value — effectively **losing one increment**.

So, the Mutex ensures that **only one goroutine** can increment `views` at a time.

---

## 🔧 Code with Mutex (Your Example)

```go
type post struct {
	views int
	mu    sync.Mutex
}

func (p *post) inc(wg *sync.WaitGroup) {
	defer func() {
		p.mu.Unlock() // Unlock after critical section
		wg.Done()     // Signal this goroutine is done
	}()
	p.mu.Lock()        // Lock before entering critical section
	p.views += 1       // Critical section: modify shared data
}
```

* `p.mu.Lock()` blocks other goroutines from entering this critical section.
* Once the critical section is done, `p.mu.Unlock()` is called using `defer` (guarantees execution even if there's an error or return).
* `wg.Done()` signals that this goroutine has finished its task.

---

## 🔁 Data Flow (Step-by-Step)

### ➕ Step-by-Step Execution:

| Step | Description                                                                                                            |
| ---- | ---------------------------------------------------------------------------------------------------------------------- |
| 1    | `main` creates a `post` instance `myPost` with `views = 0` and a `sync.WaitGroup` `wg`                                 |
| 2    | A loop starts 100 **goroutines**, each calling `myPost.inc(&wg)`                                                       |
| 3    | Each goroutine: <ul><li>Locks `mu`</li><li>Increments `views`</li><li>Unlocks `mu`</li><li>Calls `wg.Done()`</li></ul> |
| 4    | `main` waits for all goroutines using `wg.Wait()`                                                                      |
| 5    | When all goroutines are done, it prints `myPost.views` — which will now be exactly `100`                               |

---

## 🧠 Why Use `defer` for Unlock?

```go
defer func() {
    p.mu.Unlock()
    wg.Done()
}()
```

Using `defer` ensures the `Unlock` happens **no matter what** — if the function panics or returns early, you won’t leave the `Mutex` locked, which would **deadlock** the program.

---

## ⚠️ Without Mutex: What Happens?

If you remove `p.mu.Lock()` and `Unlock()`:

* You have 100 goroutines incrementing `views` simultaneously
* The increments will interfere with each other
* **Final output may be 94, 98, 99... (less than 100)** due to race conditions

---

## 📈 Visual: Concurrency with and without Mutex

### 1. Without Mutex (Race Condition)

```
Goroutine A:      views = 0 → (read)
                                ↑
Goroutine B:      views = 0 → (read)
Goroutine A:                 → views = 1 (write)
Goroutine B:                 → views = 1 (write again!)
→ Lost update: views should be 2, but is 1
```

### 2. With Mutex (No Race Condition)

```
Goroutine A: Lock  → views = 0 → views = 1 → Unlock
Goroutine B: Wait  → Lock     → views = 1 → views = 2 → Unlock
Goroutine C: Wait  → ...
→ No interference; final value = 100
```

---

## 🧮 Final Output

```go
fmt.Println(myPost.views)
```

✅ With proper Mutex use, this **always** prints:

```
100
```

---

## 📊 Diagram: Flow of Goroutines with Mutex

Here's a basic flowchart showing how the goroutines access the `views` field safely:

```
 ┌────────────┐
 │ main()     │
 └────┬───────┘
      │
      ▼
  Create myPost
      │
      ▼
  Start 100 goroutines
      │
      ▼
┌────────────┐    ┌────────────┐    ┌────────────┐
│ Goroutine 1│    │ Goroutine 2│... │ Goroutine N│
└────┬───────┘    └────┬───────┘    └────┬───────┘
     │ Lock()         │ Lock()         │ Lock()
     ▼                ▼                ▼
views += 1         views += 1       views += 1
     │                │                │
     ▼                ▼                ▼
  Unlock()         Unlock()         Unlock()
     │                │                │
     ▼                ▼                ▼
   Done()           Done()           Done()
      \              |               /
       \             |              /
        └──────────WaitGroup────────┘
                  Wait()
                   │
                   ▼
           Print views = 100 ✅
```

---

## ✅ Summary

| Concept     | Explanation                                                                         |
| ----------- | ----------------------------------------------------------------------------------- |
| Mutex       | Prevents race conditions by allowing only one goroutine to access `views` at a time |
| Lock/Unlock | Locks before modifying shared data, unlocks after done                              |
| defer       | Ensures cleanup (Unlock and Done) even if something goes wrong                      |
| WaitGroup   | Makes `main` wait for all goroutines to finish                                      |
| Benefit     | Correct and predictable results (`views = 100`)                                     |

---

