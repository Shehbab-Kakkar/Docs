 In Go, both `len` and `cap` are used with slices, but they represent **different things**.

---

### ✅ Your Code:

```go
package main
import "fmt"

func main() {
    var nums = make([]int, 2)

    fmt.Println(cap(nums)) // Capacity
    fmt.Println(len(nums)) // Length
}
```

### ➕ Output:

```
2
2
```

Yes — in this specific case, both `len(nums)` and `cap(nums)` return `2`, but **they are not always the same**.

---

## 🔍 What's the Difference?

### `len(slice)`

* Tells you **how many elements are actually in the slice**.
* It's how many items you're currently working with.

### `cap(slice)`

* Tells you **how much space is allocated** for the slice.
* It’s the total number of elements the slice can grow to **without reallocating memory**.

---

## 🧪 Example to Show the Difference:

```go
package main
import "fmt"

func main() {
    nums := make([]int, 2, 5) // length = 2, capacity = 5

    fmt.Println("len:", len(nums)) // 2
    fmt.Println("cap:", cap(nums)) // 5

    nums = append(nums, 100, 200)
    
    fmt.Println("After append:")
    fmt.Println("len:", len(nums)) // 4
    fmt.Println("cap:", cap(nums)) // 5 (still within capacity)
}
```

* Here, you're creating a slice of length 2 but capacity 5.
* You can append up to 3 more elements **without** allocating new memory.

---

## 📌 Summary

| Function | Meaning                         | Example Value |
| -------- | ------------------------------- | ------------- |
| `len()`  | Current number of elements      | `2`           |
| `cap()`  | Total allocated space for slice | `5`           |

> In your original code, since you used `make([]int, 2)`, Go sets both the length and capacity to 2. That’s why they look the same — but they aren’t always.

Let me know if you want to see how slices grow when you keep appending!
