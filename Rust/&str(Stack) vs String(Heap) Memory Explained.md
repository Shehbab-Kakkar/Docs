# Rust: `&str` (Stack) vs `String` (Heap) Memory Explained

In Rust, understanding **stack** vs **heap** memory is important for how data is stored and managed, especially with string types.

---

## 1. `&str` (String Slice) – Usually Stack Memory

- `&str` is a **string slice**: a reference to a sequence of UTF-8 bytes.
- **Stack:** `&str` itself (the reference and length) is typically stored on the stack.
- The actual string data for string literals is stored in the program's binary (read-only memory), not on heap.

**Example:**

```rust
fn main() {
    let s: &str = "hello";
    println!("{}", s);
}
```
- `"hello"` is a string literal. The data is embedded in the binary, and `s` is a reference (pointer + length) on the stack.

---

## 2. `String` – Heap Memory

- `String` is a heap-allocated, growable string type.
- **Stack:** Holds a pointer to the heap, length, and capacity.
- **Heap:** The actual contents of the string.

**Example:**

```rust
fn main() {
    let s: String = String::from("hello");
    println!("{}", s);
}
```
- `String::from("hello")` allocates memory on the heap and copies `"hello"` into it.
- The variable `s` (on the stack) holds a pointer to the heap data, its length, and capacity.

---

## 3. Visual Comparison

| Type   | Stack Stores           | Heap Stores                | Example Creation          |
|--------|------------------------|----------------------------|--------------------------|
| &str   | Pointer and length     | (Literals: binary segment) | `let s = "hi";`          |
| String | Pointer, length, cap.  | Actual string content      | `let st = String::from("hi");` |

---

## 4. Example: Both Together

```rust
fn main() {
    let stack_str: &str = "world";               // &str, reference on stack, data in binary
    let heap_string: String = String::from("hi");// String, pointer on stack, data on heap

    println!("Stack: {}", stack_str);
    println!("Heap: {}", heap_string);
}
```

---

## 5. Summary Table

| Feature        | `&str`                        | `String`                      |
|----------------|------------------------------|-------------------------------|
| Size           | Fixed, known at compile time | Growable, dynamic             |
| Memory         | Reference on stack           | Metadata on stack, data on heap|
| Use case       | For static or borrowed data  | For owned, modifiable data    |
| Example        | `"text"`                     | `String::from("text")`        |

---

**In summary:**  
- `&str` is a reference, usually points to static data, and is lightweight (stack).
- `String` is an owned, growable string stored on the heap, with metadata on the stack.
