## Primitive Data Types in Rust

Rust provides several **primitive data types** that serve as the building blocks for all programs:

---

### 1. **Scalar Types**

- **Integer**  
  - Signed: `i8`, `i16`, `i32`, `i64`, `i128`, `isize`  
  - Unsigned: `u8`, `u16`, `u32`, `u64`, `u128`, `usize`  
  - Example:  
    ```rust
    let age: u32 = 30;
    let temperature: i8 = -5;
    ```

- **Floating-Point**  
  - `f32` (32-bit), `f64` (64-bit, default)
  - Example:  
    ```rust
    let pi: f64 = 3.1415;
    let score: f32 = 99.9;
    ```

- **Boolean**  
  - `bool` — can be `true` or `false`
  - Example:  
    ```rust
    let is_rust_fun: bool = true;
    ```

- **Character**  
  - `char` — a single Unicode scalar value
  - Example:  
    ```rust
    let letter: char = 'A';
    let emoji: char = '🚀';
    ```

---

### 2. **Compound Types**

- **Tuple**  
  - Group multiple values of different types  
  - Example:
    ```rust
    let person: (u32, f64, char) = (25, 175.5, 'M');
    ```

- **Array**  
  - Fixed-size collection of values of the same type  
  - Example:
    ```rust
    let numbers: [i32; 3] = [1, 2, 3];
    ```

---

### **Summary Table**

| Type         | Example                      | Description                     |
|--------------|-----------------------------|---------------------------------|
| Integer      | `let x: i32 = 5;`           | Signed or unsigned              |
| Float        | `let y: f64 = 2.5;`         | Floating-point numbers          |
| Boolean      | `let b: bool = true;`       | `true` or `false`               |
| Character    | `let c: char = 'z';`        | Unicode character               |
| Tuple        | `let t = (1, 2.0, 'a');`    | Group of multiple values        |
| Array        | `let a = [1, 2, 3];`        | Fixed-size, same-type elements  |

---

**Tip:**  
- Use primitive types for efficient, low-level data representation and performance in Rust.
