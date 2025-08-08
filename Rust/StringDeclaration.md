## How to Declare String Type Variables in Rust

Rust provides two primary ways to declare string type variables:

---

### 1. String Slice (`&str`)

A string slice is a reference to a string, usually a string literal.

```rust
let greeting: &str = "Hello, world!";
// Or, with type inference:
let greeting = "Hello, world!";
```

- `&str` is a borrowed reference, typically immutable.

---

### 2. Owned String (`String`)

An owned `String` is a growable, heap-allocated string.

```rust
let message: String = String::from("Hello, world!");
// Or, with type inference:
let message = String::from("Hello, world!");
```

You can also convert a string slice to a `String`:

```rust
let message = "Hello, world!".to_string();
```

- `String` is useful when you need ownership or want to modify the string.

---

### Summary Table

| Type     | Example                                  | Description                     |
|----------|------------------------------------------|---------------------------------|
| `&str`   | `let s: &str = "hi";`                   | String slice (borrowed)         |
| `String` | `let s: String = String::from("hi");`    | Owned, growable heap string     |

---

**Tip:**  
- Use `&str` for string literals and when borrowing is sufficient.
- Use `String` when you need an owned, mutable, or growable string.

---

## How to Print String Values in Rust

To print variables (including string types) in Rust, use the `println!` macro.

---

### Printing a String Slice (`&str`)

```rust
let greeting = "Hello, world!";
println!("{}", greeting); // Output: Hello, world!
```

---

### Printing an Owned String (`String`)

```rust
let message = String::from("Hello, Rust!");
println!("{}", message); // Output: Hello, Rust!
```

---

### Printing Multiple Values

You can print multiple variables by using `{}` placeholders in the string:

```rust
let name = "Alice";
let language = String::from("Rust");
println!("{} loves {} programming!", name, language);
// Output: Alice loves Rust programming!
```

---

**Tip:**  
- Use `{}` as a placeholder for variables inside the `println!` string.
- List variables after the comma in the order they should be printed.
