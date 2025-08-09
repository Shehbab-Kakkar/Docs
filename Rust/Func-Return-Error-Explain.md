# Rust Implicit Return Values Explained

In Rust, functions can return a value **without** the `return` keyword. This is called an **implicit return**.

---

## Your Function Example

```rust
fn cal_len(strlen: &String) -> usize {
    strlen.len() // No semicolon! This value is returned
}
```

- If the **last line** of a function is an expression **without a semicolon**, its result is returned.

---

## What if You Add a Semicolon?

```rust
fn cal_len(strlen: &String) -> usize {
    strlen.len(); // With semicolon: this is a statement, not a return value
}
```

- With a semicolon, the last line becomes a **statement** (not an expression), so the function returns the unit type `()` by default.
- This will cause a compiler error if your function’s return type is not `()`.

### Example Error

```
error[E0308]: mismatched types
 --> main.rs:6:5
  |
6 |     strlen.len();
  |     ^^^^^^^^^^^^^ expected `usize`, found `()`
```

---

## Explicit Return (Optional)

You can always use the `return` keyword:

```rust
fn cal_len(strlen: &String) -> usize {
    return strlen.len();
}
```

But it's more idiomatic in Rust to use the implicit return by omitting the semicolon on the final expression.

---

## Summary Table

| Syntax                              | Return Value                |
|--------------------------------------|-----------------------------|
| `expr` (no semicolon)                | Value of `expr`             |
| `expr;` (with semicolon)             | `()` (unit type)            |
| `return expr;`                       | Value of `expr`             |

---

**Tip:**  
- Use no semicolon at the end of the last expression for implicit return.
- Use `return` for early or explicit returns.
