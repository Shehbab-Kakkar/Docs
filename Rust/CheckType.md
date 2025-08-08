## How to Get the Type of a Variable in Rust

Rust does not have a built-in runtime function to directly get the type of a variable, but you can use some techniques during development:

---

### 1. **Using the Compiler and Type Annotations**

If you want to check the type, you can deliberately annotate the variable with the expected type or let the compiler infer it. If there’s a type mismatch, the compiler will give an error.

---

### 2. **Using Debugging Macros**

You can use the following trick to print the type at compile time:

```rust
fn print_type_of<T>(_: &T) {
    println!("{}", std::any::type_name::<T>())
}

fn main() {
    let x = 42;
    print_type_of(&x); // Output: i32

    let s = "hello";
    print_type_of(&s); // Output: &str
}
```

- `std::any::type_name::<T>()` returns the name of the type as a string (at runtime).
- This is most useful for debugging and learning.

---

### 3. **Compiler Error Messages**

If you pass a variable to a function that expects a specific type, the compiler will tell you the type mismatch in its error messages.

---

**Summary:**  
- Use a helper like `print_type_of` with `std::any::type_name` to print the type at runtime for debugging.
- There is no built-in method to get a variable's type at runtime as you would in dynamic languages.

---
