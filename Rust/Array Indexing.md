## Array Indexing and Printing in Rust

This example demonstrates how to declare, initialize, index, and print arrays in Rust.

---

### ✅ Fixed Code Example

```rust
fn main() {
    let mut array_1: [i32; 5] = [1, 2, 3, 4, 5];
    let num = array_1[3];
    println!("{:?}", array_1);
    println!("The value at index 3 is: {}", num);
}
```

---

### 🧠 Explanation

- `let mut array_1: [i32; 5]`  
  Declares a mutable array named `array_1` that holds 5 integers.

- `[1, 2, 3, 4, 5]`  
  Initializes the array with 5 values.

- `array_1[3]`  
  Indexing starts at 0, so this accesses the 4th element (which is `4`).

- `println!("{:?}", array_1);`  
  The `{:?}` format specifier prints the whole array in debug format.

- `println!("The value at index 3 is: {}", num);`  
  Prints the specific value at index 3.

---

### 🖨️ Output

```
[1, 2, 3, 4, 5]
The value at index 3 is: 4
```

---

**Tip:**  
You can mutate (change) elements in the array since it's declared as `mut`. Let me know if you'd like an example of that!
