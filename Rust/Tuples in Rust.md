## Tuples in Rust: Example, Destructuring, and Index Access

This example demonstrates how tuples work in Rust, including destructuring and index-based access.

---

### ✅ Full Code Example

```rust
fn main() {
    let tup: (i32, f64, char) = (500, 6.4, 'a');
    let (x, y, z) = tup; // destructuring
    println!("The value of y is: {}", y);

    // Access by index
    println!("First value: {}", tup.0); // 500
    println!("Second value: {}", tup.1); // 6.4
    println!("Third value: {}", tup.2); // 'a'
}
```

---

### 🔍 Line-by-Line Explanation

- `fn main() {`  
  The entry point of the Rust program.

- `let tup: (i32, f64, char) = (500, 6.4, 'a');`  
  Declares a tuple named `tup` containing three values of different types:
  - `500` (`i32` integer)
  - `6.4` (`f64` floating-point)
  - `'a'` (`char` character)
  - The type annotation `(i32, f64, char)` is optional but shows the exact types.

- `let (x, y, z) = tup;`  
  **Destructuring**: Breaks the tuple into individual variables.
  - `x = 500`
  - `y = 6.4`
  - `z = 'a'`

- `println!("The value of y is: {}", y);`  
  Prints the value of `y` (6.4).

- **Access by index:**  
  ```rust
  println!("First value: {}", tup.0); // 500
  println!("Second value: {}", tup.1); // 6.4
  println!("Third value: {}", tup.2); // 'a'
  ```
  - Rust allows accessing tuple elements using dot notation and index:
    - `tup.0` → first element (500)
    - `tup.1` → second element (6.4)
    - `tup.2` → third element ('a')

---

### 🧠 Key Concepts

| Concept        | Meaning                                            |
|----------------|---------------------------------------------------|
| Tuple          | A fixed-size group of values of different types    |
| Destructuring  | Breaking a tuple into separate variables           |
| Index Access   | Accessing elements using `.0`, `.1`, `.2`, etc.    |

---
