# Rust Enum Example: Currency Conversion Program

This program demonstrates how to use enums, methods, pattern matching, and the Display trait in Rust to represent different currencies and their conversion rates to the US Dollar.

---

## Program Code

```rust
use std::fmt;

// Define the enum
enum Currency {
    Rupee,
    Taka,
    Yen,
    Yuan,
}

// Implement Display for Currency to allow printing with {}
impl fmt::Display for Currency {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        let name = match self {
            Currency::Rupee => "Rupee",
            Currency::Taka => "Taka",
            Currency::Yen => "Yen",
            Currency::Yuan => "Yuan",
        };
        write!(f, "{}", name)
    }
}

// Implement methods for Currency
impl Currency {
    fn currency_rate(&self) -> u32 {
        match self {
            Currency::Rupee => 87,
            Currency::Taka => 220,
            Currency::Yen => 1000,
            Currency::Yuan => 7,
        }
    }
}

fn main() {
    let current_check = Currency::Yuan;
    println!(
        "Dollar to Currency conversion for {} is {}",
        current_check,
        current_check.currency_rate()
    );
}
```

---

## Explanation

### 1. Enum Declaration

```rust
enum Currency {
    Rupee,
    Taka,
    Yen,
    Yuan,
}
```
- The `Currency` enum represents different currencies: Rupee, Taka, Yen, and Yuan.

---

### 2. Implementing the Display Trait

```rust
impl fmt::Display for Currency {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        let name = match self {
            Currency::Rupee => "Rupee",
            Currency::Taka => "Taka",
            Currency::Yen => "Yen",
            Currency::Yuan => "Yuan",
        };
        write!(f, "{}", name)
    }
}
```
- The `Display` trait allows using `{}` in `println!` macros to print the name of the currency.
- The `match` statement matches the enum variant and returns its name as a string.

---

### 3. Method for Conversion Rate

```rust
impl Currency {
    fn currency_rate(&self) -> u32 {
        match self {
            Currency::Rupee => 87,
            Currency::Taka => 220,
            Currency::Yen => 1000,
            Currency::Yuan => 7,
        }
    }
}
```
- The `currency_rate` method returns the conversion rate from US Dollar to the selected currency.
- The method takes `&self`, a reference to the current enum value, and returns a `u32` integer.

---

### 4. Main Function

```rust
fn main() {
    let current_check = Currency::Yuan;
    println!(
        "Dollar to Currency conversion for {} is {}",
        current_check,
        current_check.currency_rate()
    );
}
```
- Creates a variable `current_check` as `Currency::Yuan`.
- Prints the conversion rate for the selected currency using the `currency_rate()` method and the Display implementation.

---

## Output

```
Dollar to Currency conversion for Yuan is 7
```

---

## Key Concepts

| Concept       | Description                                                          |
|---------------|----------------------------------------------------------------------|
| Enum          | A type that can be one of several defined variants                   |
| impl          | Used to add methods (like functions) to an enum or struct            |
| Display trait | Allows pretty-printing the enum variant with `{}` in macros          |
| match         | Used to branch code based on the enum variant                        |
| Method        | `currency_rate` returns the conversion rate for the selected currency|

---

**Summary:**  
This program shows how to model multiple currencies using enums, print their names, and get their conversion rates by attaching methods to the enum in Rust.
