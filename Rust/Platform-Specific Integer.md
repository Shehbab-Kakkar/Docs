## Platform-Specific Integer Types in Rust

Rust provides two integer types whose sizes depend on the target platform:

### `isize` and `usize`

- **`isize`**: A signed integer type, with a size equal to the pointer size of the platform (e.g., 32 bits on 32-bit systems, 64 bits on 64-bit systems).
- **`usize`**: An unsigned integer type, also matching the platform’s pointer size.

#### **Typical Use Cases**
- Used for indexing collections (like arrays, vectors).
- Useful when dealing with memory addresses or sizes.

#### **Examples**

```rust
let signed: isize = -123;
let unsigned: usize = 456;
println!("signed isize: {}", signed);
println!("unsigned usize: {}", unsigned);
```

#### **Why Use `isize` and `usize`?**
- They ensure your code works efficiently and correctly on both 32-bit and 64-bit systems, especially for operations involving memory or collection indices.

---

**Summary Table**

| Type    | Description                      | Example                  |
|---------|----------------------------------|--------------------------|
| `isize` | Signed, platform pointer-sized   | `let a: isize = -10;`    |
| `usize` | Unsigned, platform pointer-sized | `let b: usize = 20;`     |
