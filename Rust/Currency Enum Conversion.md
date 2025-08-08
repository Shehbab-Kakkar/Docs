# Detailed Explanation: Currency Enum Conversion Program in Rust

This program demonstrates how to use enums, methods, pattern matching, and formatting in Rust to represent different world currencies and their conversion rates with respect to the US Dollar.

---

## Complete Code

```rust
enum Currency {
    Rupee,
    Taka,
    Yen,
    Yuan,
}

impl Currency {
    fn currency_rate(&self) -> u32 {
        match self {
            Currency::Rupee => 87,
            Currency::Taka => 220,
            Currency::Yen => 1000,
            Currency::Yuan => 7,
        }
    }
    
    fn as_str(&self) -> &str {
        match self {
            Currency::Rupee => "Rupee",
            Currency::Taka => "Taka",
            Currency::Yen => "Yen",
            Currency::Yuan => "Yuan",
        }
    }
}

fn main() {
    let current_check = Currency::Yen;
    println!(
        "Dollar to Currency conversion for {} is {}",
        current_check.as_str(),
        current_check.currency_rate()
    );
}
```

---

## Keyword and Data Flow Breakdown

### 1. `enum`  
- **Keyword:** Defines an enumeration, a type which can be one of several named values (called variants).
- **Here:**  
  ```rust
  enum Currency {
      Rupee,
      Taka,
      Yen,
      Yuan,
  }
  ```
  This creates a data type called `Currency` with four possible values: `Rupee`, `Taka`, `Yen`, and `Yuan`.

---

### 2. `impl`  
- **Keyword:** Used to implement methods (functions) for a type (here, for `Currency` enum).
- **Here:**  
  ```rust
  impl Currency {
      ...
  }
  ```
  All functions inside this block are available to every `Currency` value.

---

### 3. `fn currency_rate(&self) -> u32`
- **`fn`:** Defines a function/method.
- **`currency_rate`:** Name of the method.
- **`&self`:** Refers to the instance of `Currency` on which the method is called.
- **`-> u32`:** The return type; returns an unsigned 32-bit integer.

**Functionality:**  
Uses a `match` statement to return the exchange rate for the specific currency variant.

---

### 4. `match self`
- **Keyword:** Pattern matching. Branches code based on the value of `self` (the current enum variant).
- **Here:**  
  ```rust
  match self {
      Currency::Rupee => 87,
      Currency::Taka => 220,
      Currency::Yen => 1000,
      Currency::Yuan => 7,
  }
  ```
  Returns an integer based on which `Currency` is used.

---

### 5. `fn as_str(&self) -> &str`
- **Purpose:** Returns the name of the currency as a string slice (`&str`).  
- **Why:** This allows pretty-printing the enum name in output.

---

### 6. `fn main()`
- **Main entry point** for every Rust program.
- **Here:**  
  ```rust
  let current_check = Currency::Yen;
  ```
  - Creates a variable named `current_check` and assigns it the enum variant `Currency::Yen`.

---

### 7. `println!` macro
- Used for printing to the console.  
- Supports formatting with `{}` placeholders.  
- **Here:**  
  ```rust
  println!(
      "Dollar to Currency conversion for {} is {}",
      current_check.as_str(),
      current_check.currency_rate()
  );
  ```
  - The first `{}` is replaced by the result of `current_check.as_str()` ("Yen")
  - The second `{}` is replaced by the result of `current_check.currency_rate()` (1000)

---

## Data Flow

1. **Program starts in `main()`.**
2. `current_check` is set to `Currency::Yen`.
3. `current_check.as_str()` returns `"Yen"` for display.
4. `current_check.currency_rate()` returns `1000` (exchange rate).
5. Both values are printed in the formatted string:  
   ```
   Dollar to Currency conversion for Yen is 1000
   ```

---

## Key Concepts Table

| Keyword     | Meaning                                                                | Example in Program                  |
|-------------|-----------------------------------------------------------------------|-------------------------------------|
| `enum`      | Custom type with named variants                                       | `enum Currency { ... }`             |
| `impl`      | Implementation block for methods                                      | `impl Currency { ... }`             |
| `fn`        | Function/method definition                                            | `fn as_str(&self) -> &str`          |
| `&self`     | Reference to current instance (immutable)                             | Method receivers                    |
| `match`     | Pattern matching, selects code based on enum value                    | `match self { ... }`                |
| `let`       | Variable binding                                                      | `let current_check = Currency::Yen` |
| `println!`  | Macro for formatted output                                            | `println!(...)`                     |
| `-> u32`    | Function return type (unsigned 32-bit integer)                        | `fn currency_rate(&self) -> u32`    |
| `-> &str`   | Function return type (string slice)                                   | `fn as_str(&self) -> &str`          |

---

## Summary

- **Enums** let you define a type with a fixed set of possible values.
- **Methods** can be attached to enums using `impl`.
- **Pattern matching** with `match` lets you run code depending on which enum variant you have.
- **Data flows** from variable assignment, through method calls, into formatted output.

This program is a clear example of Rust's type safety, pattern matching, and structured data handling!
