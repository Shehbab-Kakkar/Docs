# Rust Ownership Explained

**Ownership** is one of Rust’s core concepts for managing memory safely and efficiently without a garbage collector. It ensures safety by enforcing rules at compile time.

## Key Rules of Ownership

1. **Each value in Rust has a single owner.**
2. **When the owner goes out of scope, the value is dropped and memory freed.**
3. **A value can be moved (transferred) to a new owner, but only one owner at a time.**

---

## Example: Ownership in Action

```rust
fn main() {
    let s1 = String::from("hello");
    let s2 = s1; // Ownership of the String is moved from s1 to s2

    // println!("{}", s1); // This would cause a compile error: value borrowed after move
    println!("{}", s2); // This works: s2 is now the owner
}
```

### What Happens Here?
- `let s1 = String::from("hello");`  
  `s1` owns the String `"hello"`.
- `let s2 = s1;`  
  Ownership moves from `s1` to `s2`. Now `s1` is **invalid**.
- Trying to use `s1` after the move causes a **compile-time error**.

---

## Borrowing Example

You can **borrow** a value using references:

```rust
fn main() {
    let s1 = String::from("hello");
    let len = calculate_length(&s1); // Pass reference, not ownership
    println!("The length of '{}' is {}.", s1, len); // s1 is still valid
}

fn calculate_length(s: &String) -> usize {
    s.len()
}
```

- `&s1` passes a **reference** of `s1`, not its ownership.
- `s1` can still be used after the function call.

---

## Mutable Borrow Example

```rust
fn main() {
    let mut s = String::from("hello");
    change(&mut s);
    println!("{}", s); // prints "hello, world"
}

fn change(some_string: &mut String) {
    some_string.push_str(", world");
}
```
- `&mut s` borrows `s` as mutable, allowing the function to modify it.

---

## Summary Table

| Concept          | What it means                             | Example                                    |
|------------------|-------------------------------------------|--------------------------------------------|
| Ownership Move   | Ownership is transferred                  | `let s2 = s1;`                             |
| Borrowing        | Reference, not ownership, is passed       | `calculate_length(&s1)`                    |
| Mutable Borrow   | Mutable reference for safe modification   | `change(&mut s)`                           |

---

**In summary:**  
Ownership ensures memory safety in Rust by allowing only one owner at a time for each value, enforcing clean-up when the owner is out of scope, and using borrowing for safe references.
