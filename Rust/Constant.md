## Constants in Rust

In Rust, you can define constants using the `const` keyword. Constants are:

- Always immutable (cannot be changed after declaration).
- Must have an explicit type annotation.
- Evaluated at compile time.
- Typically written in uppercase with underscores (by convention).
- Can be declared in any scope, including global.

### Syntax

```rust
const MAX_POINTS: u32 = 100;
```

### Example

```rust
const PI: f64 = 3.14159;

fn main() {
    println!("The value of PI is: {}", PI);
}
```

### Key Points

- Constants cannot be assigned the result of a function call that is not a `const fn` or any value that is not known at compile time.
- Use `const` for values that should never change and are known at compile time.

---

**Tip:**  
If you need a value that can be set at runtime and won't change afterward, consider using `let` with the `mut` keyword (for mutability) or without it (for immutability).
