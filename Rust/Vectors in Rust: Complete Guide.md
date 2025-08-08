# Vectors in Rust: Complete Guide

Vectors (`Vec<T>`) are the most commonly used **growable, heap-allocated, resizable arrays** in Rust. Unlike fixed-size arrays (`[T; N]`), vectors can dynamically change their size at runtime.

---

## Declaring and Initializing a Vector

You can create a new, empty vector using `Vec::new()`:

```rust
let mut vector_name: Vec<i32> = Vec::new();
```

Or using the macro:

```rust
let mut vector_name = vec![1, 2, 3];
```

---

## Adding Elements

- **push(value):** Adds an element to the end of the vector.

    ```rust
    vector_name.push(10);
    vector_name.push(20);
    vector_name.push(30);
    ```

---

## Removing Elements

- **pop():** Removes the last element and returns it as an `Option<T>`. Returns `None` if the vector is empty.

    ```rust
    let last = vector_name.pop(); // Some(30)
    ```

- **remove(index):** Removes and returns the element at the given index, shifting all elements after it to the left.

    ```rust
    vector_name.remove(1); // Removes the element at index 1 (second element)
    ```

---

## Accessing Elements

- **Indexing:** Like arrays, you can access elements using square brackets.

    ```rust
    let first = vector_name[0];
    ```

- **get(index):** Returns an `Option<&T>`, which is safer than direct indexing (avoids panic if out of bounds).

    ```rust
    if let Some(value) = vector_name.get(2) {
        println!("Value: {}", value);
    }
    ```

---

## Modifying and Other Operations

- **reverse():** Reverses the order of elements in the vector.

    ```rust
    vector_name.reverse();
    ```

- **len():** Returns the number of elements in the vector.

    ```rust
    let size = vector_name.len();
    ```

- **is_empty():** Checks if the vector is empty.

    ```rust
    if vector_name.is_empty() {
        println!("Vector is empty!");
    }
    ```

- **clear():** Removes all elements from the vector.

    ```rust
    vector_name.clear();
    ```

---

## Complete Example

```rust
fn main() {
    let mut vector_name: Vec<i32> = Vec::new();
    vector_name.push(10);
    vector_name.push(20);
    vector_name.push(30);

    println!("{:?}", vector_name); // [10, 20, 30]

    // vector_name.pop();         // Removes last element (30)
    vector_name.remove(1);        // Removes element at index 1 (20), now [10, 30]
    vector_name.reverse();        // Now [30, 10]

    println!("{:?}", vector_name); // [30, 10]
}
```

---

## Output

```
[10, 20, 30]
[30, 10]
```

---

## Important Vector Functions (Summary Table)

| Function        | Description                                      | Example                         |
|-----------------|--------------------------------------------------|---------------------------------|
| `push(value)`   | Add value to end                                 | `v.push(5)`                     |
| `pop()`         | Remove and return last value (Option)            | `let x = v.pop()`               |
| `remove(index)` | Remove and return value at given index           | `v.remove(1)`                   |
| `get(index)`    | Get value as Option<&T>                          | `v.get(2)`                      |
| `len()`         | Get number of elements                           | `v.len()`                       |
| `is_empty()`    | Returns true if vector is empty                  | `v.is_empty()`                  |
| `reverse()`     | Reverse the vector in place                      | `v.reverse()`                   |
| `clear()`       | Remove all elements                              | `v.clear()`                     |

---

## Key Points

- Vectors are **heap-allocated**, so their size can grow or shrink at runtime.
- Always use `mut` if you want to modify a vector after creation.
- Vectors are **zero-indexed**: the first element is at index `0`.
- Vectors are the go-to choice for most dynamic list needs in Rust.

---

**Tip:**  
For more information, see [Rust's Vec documentation](https://doc.rust-lang.org/std/vec/struct.Vec.html).
