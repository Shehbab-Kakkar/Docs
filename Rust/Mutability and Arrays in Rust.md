## Mutability and Arrays in Rust

Rust arrays have a **fixed size**, but whether their contents can be changed depends on how they are declared.

---

### Declaring a Mutable Array

```rust
let mut array_1: [i32; 5] = [1, 2, 3, 4, 5];
```

- `mut` means **you can change the values inside the array**.
- The **size** (`[i32; 5]`) is fixed at compile time and cannot be changed.
- Example:
  ```rust
  array_1[0] = 10; // OK!
  // array_1.push(6); // ERROR! Arrays can't change size.
  ```

---

### Declaring an Immutable Array

```rust
let array_1: [i32; 5] = [1, 2, 3, 4, 5];
```

- **No `mut`** means the array is **immutable**.
- You **cannot change any element** after creation.
- Example:
  ```rust
  array_1[0] = 10; // ERROR! Cannot assign to `array_1[0]` because it is not mutable
  ```

---

### Why Rust Enforces Immutability by Default

- **Safer code**: Prevents accidental changes.
- **Predictability**: Easier to reason about program state.
- **Performance**: Compiler can optimize better.

---

### Resizable Collections: Use a Vector (`Vec<T>`)

If you want a resizable collection, use a `Vec<T>`:

```rust
let mut vec_1 = vec![1, 2, 3, 4, 5];
vec_1.push(6); // Works! Vec is growable
println!("{:?}", vec_1); // [1, 2, 3, 4, 5, 6]
```

---

### ✨ Summary Table

| Declaration                 | Mutability       | Can modify elements? | Can resize?       |
|-----------------------------|------------------|----------------------|-------------------|
| `let array_1 = [...]`       | Immutable        | ❌ No                | ❌ No             |
| `let mut array_1 = [...]`   | Mutable          | ✅ Yes               | ❌ No             |
| `let mut vec_1 = vec![...]` | Mutable (Vector) | ✅ Yes               | ✅ Yes (growable) |

---

**TL;DR:**  
- `mut` on an array lets you change values, **not the size**.
- Use `Vec<T>` for a collection you can resize.

---
