## Example: Creating an Empty Vector in Rust

You can create an empty vector using `Vec::new()` or the `vec![]` macro.

### Using `Vec::new()`:
```rust
let mut v: Vec<i32> = Vec::new();
```
- This creates an empty, mutable vector of `i32`.
- You can then add elements using `.push()`.

### Using the `vec![]` macro:
```rust
let mut v: Vec<i32> = vec![];
```
- This also creates an empty, mutable vector of `i32`.

### Example: Adding Elements Later
```rust
let mut v: Vec<i32> = Vec::new();
v.push(10);
v.push(20);
println!("{:?}", v); // Output: [10, 20]
```

---

**Note:**  
- You can specify the type explicitly (e.g., `Vec<i32>`) or let Rust infer it from usage.
