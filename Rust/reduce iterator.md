Great — you're now asking about **using `into_iter()` on a `Vec`**, applying **`reduce`**, and comparing that to using **`sum()`** to collect or aggregate values.

Let’s break this into clear sections with examples.

---

## ✅ 1. `into_iter()` + `reduce()`

The `.reduce()` method on an iterator **combines elements** using a closure, like folding them together. It returns an `Option<T>` because the iterator might be empty.

```rust
fn main() {
    let nums = vec![1, 2, 3, 4, 5];

    let reduced = nums
        .into_iter() // Takes ownership, yields T
        .reduce(|acc, x| acc + x);

    println!("Reduced sum: {:?}", reduced);
}
```

### 🧾 Output:

```
Reduced sum: Some(15)
```

If the vector was empty, `.reduce()` would return `None`.

---

## ✅ 2. `into_iter()` + `sum()`

The `.sum()` method **sums up all elements**, but unlike `reduce()`, it does **not** return an `Option` — it just returns `0` for empty iterators.

```rust
fn main() {
    let nums = vec![1, 2, 3, 4, 5];

    let total: i32 = nums.into_iter().sum();

    println!("Summed total: {}", total);
}
```

### 🧾 Output:

```
Summed total: 15
```

---

## 🔍 Comparison Table

| Method      | Type Signature                          | Empty Vec Result | Return Type      |
| ----------- | --------------------------------------- | ---------------- | ---------------- |
| `.reduce()` | `FnMut(T, T) -> T`                      | `None`           | `Option<T>`      |
| `.sum()`    | `Sum<T>` (requires `T: std::iter::Sum`) | `0`              | `T` (e.g. `i32`) |

---

## ✅ 3. Example: Sum Only Even Numbers with `filter`, `map`, `sum`

```rust
fn main() {
    let nums = vec![1, 2, 3, 4, 5, 6];

    let even_sum: i32 = nums
        .into_iter()
        .filter(|x| x % 2 == 0) // keep even numbers
        .map(|x| x * 2)         // double them
        .sum();                 // sum the result

    println!("Even doubled sum: {}", even_sum);
}
```

### 🧾 Output:

```
Even doubled sum: 24
```

* Even numbers: `[2, 4, 6]`
* Doubled: `[4, 8, 12]`
* Sum: `24`

---

## 🧠 Summary

| Use Case               | Method                             | Output Type | Use When...                                |
| ---------------------- | ---------------------------------- | ----------- | ------------------------------------------ |
| Total of numbers       | `.sum()`                           | `T`         | Simple summing                             |
| Custom aggregation     | `.reduce()`                        | `Option<T>` | You want custom behavior / check for empty |
| With filtering/mapping | Combine with `.filter()`, `.map()` | —           | Only include certain values                |

---

Would you like to see how `.fold()` compares with `.reduce()` too? It's similar but more flexible.
