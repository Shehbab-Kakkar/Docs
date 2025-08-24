

In Rust, the **main data type that guarantees unique elements** is:

---

## 🟢 **`HashSet`**

### ✅ Characteristics:

| Feature                       | Description                                      |
| ----------------------------- | ------------------------------------------------ |
| **Unique elements?**          | ✅ Yes — duplicates are automatically removed     |
| **Ordered?**                  | ❌ No — unordered collection                      |
| **Fixed size?**               | ❌ No — growable                                  |
| **Can hold different types?** | ❌ No — all elements must be of the **same type** |

### 🔸 Example:

```rust
use std::collections::HashSet;

fn main() {
    let mut set = HashSet::new();

    set.insert("apple");
    set.insert("banana");
    set.insert("apple"); // Duplicate — will not be added

    println!("{:?}", set); // Output: {"banana", "apple"} (order not guaranteed)
}
```

---

## 🟢 **Other Unique-Element Collections in Rust**

### 1. **`BTreeSet`**

* Like `HashSet`, but keeps elements **sorted**.
* Uses a tree structure instead of a hash table.

```rust
use std::collections::BTreeSet;

fn main() {
    let mut set = BTreeSet::new();
    set.insert(3);
    set.insert(1);
    set.insert(2);
    set.insert(1); // duplicate

    for value in &set {
        println!("{}", value); // Output: 1, 2, 3 (in order)
    }
}
```

---

## ❌ Types That Do *Not* Enforce Uniqueness

| Type      | Allows Duplicates?                            |
| --------- | --------------------------------------------- |
| `Vec<T>`  | ✅ Yes                                         |
| `Tuple`   | ✅ Yes                                         |
| `Array`   | ✅ Yes                                         |
| `String`  | ✅ Yes (characters)                            |
| `HashMap` | ✅ (keys must be unique, values can duplicate) |

---

## ✅ Summary

| Data Type     | Unique Elements | Ordered | Same Type Required |
| ------------- | --------------- | ------- | ------------------ |
| `HashSet<T>`  | ✅ Yes           | ❌ No    | ✅ Yes              |
| `BTreeSet<T>` | ✅ Yes           | ✅ Yes   | ✅ Yes              |
| `Vec<T>`      | ❌ No            | ✅ Yes   | ✅ Yes              |
| `Tuple`       | ❌ No            | ✅ Yes   | ❌ No               |

---

Let me know if you want to see how to enforce uniqueness in a `Vec`, or how to store custom types in a `HashSet`!
