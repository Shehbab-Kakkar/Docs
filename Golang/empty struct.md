 how you can declare an **empty struct** in Go and print its value.

---

### ✅ Example Code

```go
package main

import "fmt"

func main() {
	// Declare an empty struct
	var s struct{}

	// Print the value
	fmt.Println("Empty struct value:", s)
}
```

---

### 🧾 Output

```bash
Empty struct value: {}
```

---

### 📌 Explanation

* `struct{}` is an **empty struct type** in Go — it has **zero fields**, takes **zero bytes of storage**, and is commonly used when you don’t need to store any data.
* When printed, Go shows its value as `{}` (empty braces).

---

### ⚡️ Bonus Tip

Empty structs are useful in Go for things like:

* **Signaling only presence** (e.g., in sets or channels).
* **Memory-efficient markers** (since they occupy 0 bytes).

Example using `map[string]struct{}` as a set:

```go
mySet := map[string]struct{}{
    "apple": {},
    "banana": {},
}
```


