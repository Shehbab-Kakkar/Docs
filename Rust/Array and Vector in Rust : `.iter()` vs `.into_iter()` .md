# Array and Vector in Rust: `.iter()` vs `.into_iter()`

This section explains how to use `.iter()` and `.into_iter()` with arrays and vectors in Rust, and when to use each.

---

## Array

```rust
fn main() {
    let arr = [10, 20, 30, 40, 50];

    // .iter() borrows each element
    for x in arr.iter() {
        println!("iter: {}", x);
    }

    // .into_iter() moves or copies each element out (since Rust 1.53+ for all array lengths)
    for x in arr.into_iter() {
        println!("into_iter: {}", x);
    }
}
```

- **`.iter()`**: Returns references (`&i32`). The array is **not consumed** and can be used after the loop.
- **`.into_iter()`**: Returns values (`i32` for Copy types). The array is **consumed**, but for `Copy` types this is just a copy, so you can still use the original array after.

---

## Vector

```rust
fn main() {
    let vec = vec![10, 20, 30, 40, 50];

    // .iter() borrows each element
    for x in vec.iter() {
        println!("iter: {}", x);
    }

    // .into_iter() moves each element out, consuming the vector
    let vec2 = vec![10, 20, 30, 40, 50];
    for x in vec2.into_iter() {
        println!("into_iter: {}", x);
    }
    // vec2 cannot be used here anymore!
}
```

- **`.iter()`**: Borrows each element (`&i32`). The vector can still be used after.
- **`.into_iter()`**: **Consumes** the vector and moves out each element (`i32`). The vector **cannot** be used after.

---

## When to Use `.iter()` vs `.into_iter()`

- Use **`.iter()`** when you want to **keep using** the collection after the loop, or you only need references to the elements.
- Use **`.into_iter()`** when you want to **consume** the collection (often for ownership or when you want to transform or collect into another type).

---

## Summary Table

| Collection | `.iter()`               | `.into_iter()`                   |
|------------|------------------------|----------------------------------|
| **Array**  | Borrows (`&T`)         | Moves/copies (`T`), not consumed |
| **Vector** | Borrows (`&T`)         | Moves (`T`), vector is consumed  |

---

### Practical Tip
- Most of the time, use `.iter()` unless you specifically want to consume the collection.
- For iterating by reference, `.iter()` is best.
- For transforming or collecting into a new collection, `.into_iter()` is often handy.

---
