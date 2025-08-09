# Rust Example: Borrowing, Function Return Types, and String Length

## Problematic Code and Error Explanation

```rust
fn main() {
    let s1 = String::from("Borrowing");
    let s2 = cal_len(&s1); // Passing reference, not ownership
    println!("s1 = {} s2 = {}", s1, s2);
}

fn cal_len(strlen: &String) -> u8 {
    strlen.len();
}
```

### What's the Error?

- The function signature says it returns `u8`, but there is **no return value** in `cal_len`.
- `strlen.len();` computes the length but doesn't return it!
- Rust expects a `u8`, but the function returns `()` (the unit type).

#### Example Error Message:

```
error[E0308]: mismatched types
 --> src/main.rs:7:24
  |
7 | fn cal_len(strlen:&String) -> u8 {
  |                            ^^^^^ expected `u8`, found `()`
```

---

## How to Fix

- Return the length value, cast to `u8` (since `len()` returns `usize`).
- Either use a `return` statement, or just make the last expression the value.

### Corrected Code

```rust
fn main() {
    let s1 = String::from("Borrowing");
    let s2 = cal_len(&s1); // Passing reference of s1
    println!("s1 = {} s2 = {}", s1, s2);
}

fn cal_len(strlen: &String) -> u8 {
    strlen.len() as u8
}
```

---

## Key Points

- `&s1` passes a reference (borrow), so `s1` is still valid after the function call.
- `String::len()` returns a `usize`, so you need to cast it to `u8` for the function signature.
- If your string could be longer than 255 bytes, consider using `usize` or `u32` instead of `u8` to avoid truncation.

---
