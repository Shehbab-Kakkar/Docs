## Type Conversion Example in Rust Using `as`

The most common way to convert between primitive types in Rust is using the `as` keyword.

---

### ✅ Example: Converting `f64` to `u32`

```rust
fn main() {
    let pi: f64 = 3.14159;

    // Convert from f64 to u32 (truncates the decimal part)
    let integer_pi: u32 = pi as u32;

    println!("pi as f64: {}", pi);
    println!("pi as u32: {}", integer_pi);
}
```

**🟢 Output:**
```
pi as f64: 3.14159
pi as u32: 3
```

---

### 💡 Notes

- `as` performs explicit type conversion.
- Converting from `f64` to `u32` truncates the decimal (does not round).
- Be cautious when converting between signed and unsigned types or across different sizes—Rust won’t stop you at compile time, but can produce unexpected results at runtime.

---

### 🔄 More Examples of Type Conversion

**1. `i32` to `f64`**
```rust
let x: i32 = 10;
let y: f64 = x as f64;
```

**2. `u8` to `char`**
```rust
let byte: u8 = 65;
let letter: char = byte as char; // 'A'
```

**3. `usize` to `i32` (with care!)**
```rust
let index: usize = 5;
let idx_i32: i32 = index as i32;
```
