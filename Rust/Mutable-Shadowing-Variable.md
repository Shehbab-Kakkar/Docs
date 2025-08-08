## Difference Between Mutability and Shadowing in Rust

Rust provides two ways to change what a variable name refers to: **mutability** and **shadowing**. Here’s how they differ:

---

### 1. Mutability (`mut`)

- Use `mut` to make a variable's value changeable.
- **The type of the variable cannot change.**

```rust
let mut x = 5;
println!("x is: {}", x); // prints: x is: 5
x = 10; // OK, x is mutable
println!("x is: {}", x); // prints: x is: 10
// x = "hello"; // ❌ Error: cannot assign a &str to an i32
```

---

### 2. Shadowing (`let`)

- Shadowing redeclares a variable with the same name, possibly changing its type or value.
- Each `let` creates a new variable, even if the name is the same.

```rust
let y = 5;
println!("y is: {}", y); // prints: y is: 5
let y = "hello"; // shadows the previous y, type changes to &str
println!("y is: {}", y); // prints: y is: hello
```

---

### **Summary Table**

| Feature     | Mutability (`mut`)           | Shadowing (`let`)       |
|-------------|------------------------------|-------------------------|
| Type change | Not allowed                  | Allowed                 |
| Scope       | Same variable, same binding  | New variable, new binding|
| Usage       | `let mut x = ...; x = ...;`  | `let x = ...; let x = ...;` |

---

### **Combined Example**

```rust
// Mutability
let mut count = 1;
count = 2; // OK
// count = "hello"; // ❌ Error

// Shadowing
let count = 1;
let count = "hello"; // OK, type changes from i32 to &str
```

> **In short:**  
> - Use `mut` to change the value of a variable (but not its type).  
> - Use shadowing (`let`) to create a new variable with the same name, possibly with a different type or value.
