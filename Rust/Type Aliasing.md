## Type Aliasing in Rust

In Rust, **type aliasing** allows you to create a new name for an existing type using the `type` keyword. This can make complex types easier to read and maintain, or improve code clarity.

---

### **Syntax**

```rust
type AliasName = ExistingType;
```

---

### **Examples**

#### 1. Simple Type Alias

```rust
type Kilometers = i32;

let distance: Kilometers = 100;
println!("Distance: {} km", distance);
```

#### 2. Aliasing Complex Types

```rust
type Thunk = Box<dyn Fn() + Send + 'static>;

fn take_long_type(f: Thunk) {
    // ...
}
```

#### 3. Aliasing for Readability

```rust
type Score = u32;
let high_score: Score = 9000;
```

---

### **Key Points**
- Type aliases **do not create new types**, just new names for existing types.
- Useful for making signatures of functions or types involving generics or traits more readable.
- Commonly used for clarity, especially with complex types.

---

**Tip:** Use type aliases to simplify long or complicated type signatures and improve code readability.
