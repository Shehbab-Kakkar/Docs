# Rust Enum Example: Coin Value Program

This program demonstrates how to use enums and pattern matching in Rust to represent and handle different coin types, each with its own value in cents.

---

## Program Code

```rust
enum Coin {
    Penny,
    Nickel,
    Dime,
    Quarter,
}

impl Coin {
    fn value_in_cents(&self) -> u8 {
        match self {
            Coin::Penny => 1,
            Coin::Nickel => 5,
            Coin::Dime => 10,
            Coin::Quarter => 25,
        }
    }
}

fn main() {
    let coin = Coin::Dime;
    println!("The value of the coin is {} cents", coin.value_in_cents());
}
```

---

## Explanation

### 1. Enum Declaration

```rust
enum Coin {
    Penny,
    Nickel,
    Dime,
    Quarter,
}
```
- The `Coin` enum represents four US coin types: Penny, Nickel, Dime, and Quarter.
- Each variant stands for a specific coin.

---

### 2. Implementing Methods for the Enum

```rust
impl Coin {
    fn value_in_cents(&self) -> u8 {
        match self {
            Coin::Penny => 1,
            Coin::Nickel => 5,
            Coin::Dime => 10,
            Coin::Quarter => 25,
        }
    }
}
```
- The `impl` block adds methods to the `Coin` enum.
- `value_in_cents` is a method that returns the value of the coin in cents.
- It uses a `match` statement to check which variant of `Coin` is being used and returns the corresponding cent value.

---

### 3. Using the Enum and Its Method

```rust
fn main() {
    let coin = Coin::Dime;
    println!("The value of the coin is {} cents", coin.value_in_cents());
}
```
- In the `main` function, a variable `coin` is created and assigned the value `Coin::Dime`.
- It then prints the value of the coin in cents by calling the `value_in_cents` method.

---

## Output

```
The value of the coin is 10 cents
```

---

## Key Concepts

| Concept          | Description                                                           |
|------------------|-----------------------------------------------------------------------|
| Enum             | A type that can be one of several defined variants.                   |
| impl             | Used to add methods (like functions) to a type such as an enum.       |
| match            | Pattern matching, used to branch code based on enum variants.         |
| Method Call      | `coin.value_in_cents()` calls the method for that specific enum value.|

---

**In summary:**  
This program shows how to use enums to represent different types of coins and how to associate behavior (methods) with enum types using pattern matching in Rust.
