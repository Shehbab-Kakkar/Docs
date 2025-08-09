# Explanation: Stack vs Heap, Ownership, and `.clone()` in Rust

Let's break down your `main` function and understand what happens with memory for both `&str` (stack) and `String` (heap), especially focusing on the use of `.clone()`.

---

## Code

```rust
fn main(){
    // Stack String Literal
    let s1: &str = "Hello I am Stack";
    let s2 = s1;
    println!("s1 = {}, S2 = {} ", s1, s2);

    // Heap Memory String
    let s3: String = String::from("Hello I am Heap");
    let s4 = s3.clone();
    println!("Print s3 {}", s3); // This works, no error
    // Ownership Transfer from s3 to s4. Use clone() to create copy
    // And use new memory space for s4
    println!("Print s4 {}", s4);
}
```

---

## Line-by-line Explanation

### 1. Stack String Literal

```rust
let s1: &str = "Hello I am Stack";
let s2 = s1;
println!("s1 = {}, S2 = {} ", s1, s2);
```
- `"Hello I am Stack"` is a **string literal**. Its data is embedded in the program binary (read-only memory).
- `s1` is a reference (`&str`) to this static string data. The reference (pointer + length) lives on the **stack**.
- `let s2 = s1;` copies the reference (pointer and length) — both `s1` and `s2` now point to the same data; **no new memory is allocated**.
- Both `s1` and `s2` can be used safely; they do not own the data, they just refer to it.

---

### 2. Heap Memory String

```rust
let s3: String = String::from("Hello I am Heap");
let s4 = s3.clone();
println!("Print s3 {}", s3); // This works, no error
println!("Print s4 {}", s4);
```
- `String::from("Hello I am Heap")` creates a new **String** object.  
  - The `String`'s actual content (`"Hello I am Heap"`) is stored on the **heap** (dynamically allocated memory).
  - The `String` variable (`s3`) itself (pointer, length, capacity) is on the **stack**, but it "owns" the heap data.
- `let s4 = s3.clone();` creates a **deep copy**:
  - Allocates **new heap memory** for `s4`.
  - Copies the content from `s3`'s heap data into `s4`'s heap space.
  - Now, `s3` and `s4` are **independent**, each owning its own heap allocation.
- Both `s3` and `s4` are valid and can be printed, because they are two separate owners with different heap memory.

---

## `.clone()` and Memory

- `.clone()` performs a **deep copy**.
  - For `String`, this means allocating new heap memory and copying the bytes.
- Without `.clone()`, `let s4 = s3;` would **move** ownership from `s3` to `s4`.  
  - After a move, `s3` is **invalid** and cannot be used.
- Using `.clone()` allows you to keep both `s3` and `s4` valid and independent.

---

## Summary Table

| Variable | Type    | Memory Location     | Ownership         | Can use after assignment? | Allocates new memory? |
|----------|---------|--------------------|-------------------|--------------------------|----------------------|
| s1, s2   | &str    | Stack (reference)  | No (borrow only)  | Yes                      | No                   |
| s3       | String  | Stack (meta), Heap | Yes (owner)       | Yes (if cloned)          | Yes (for heap data)  |
| s4       | String  | Stack (meta), Heap | Yes (owner)       | Yes                      | Yes (for heap data)  |

---

**Key Takeaway:**  
- `&str` copies references, no new memory.
- `String` is heap-allocated and moves ownership by default.
- Use `.clone()` to explicitly copy heap data, so both variables are valid and own separate data.
