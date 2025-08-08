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




---

## How Is `value_in_cents` Called from `impl Coin`? What’s the Purpose?

In the given code:

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

### How can we call `value_in_cents`?

- The function `value_in_cents` is defined inside an `impl Coin` block, which means it is **a method for the `Coin` enum**.
- You can call this method on any instance of `Coin` (like `coin.value_in_cents()`).
- In `main`, `let coin = Coin::Dime;` creates a variable of type `Coin`, and then `coin.value_in_cents()` calls the method.

---

### What is the Purpose of `impl Coin`?

- The `impl` block is used to **implement methods** for a type (enum, struct, etc.).
- Here, we add the method `value_in_cents` to `Coin`, so every `Coin` value can use this method to get its value in cents.
- This is similar to attaching functions to a type in object-oriented languages.

---

### What does `&self` mean?

- In the method `fn value_in_cents(&self) -> u8`, the `&self` is a reference to the current instance (the `Coin` value on which the method is called).
- It's similar to `this` in other languages (like Java, C++, Python).
- The `&` means it is an **immutable reference**; the method can read but not change the value.

---

### What does `match self` mean?

- `match self` is a pattern matching expression based on the value of `self` (the current `Coin`).
- It checks which variant (`Penny`, `Nickel`, `Dime`, or `Quarter`) the current coin is and returns the corresponding value.

#### Example inside the method:
```rust
match self {
    Coin::Penny => 1,
    Coin::Nickel => 5,
    Coin::Dime => 10,
    Coin::Quarter => 25,
}
```
- If `self` is `Coin::Dime`, it matches the third arm and returns `10`.

---

### Summary Table

| Concept           | Meaning                                                                                   |
|-------------------|------------------------------------------------------------------------------------------|
| `impl Coin`       | Attach methods (functions) to the `Coin` enum                                            |
| `&self`           | Reference to the instance of `Coin` the method is called on                              |
| `match self`      | Pattern match against the current value (`self`) to decide what code to run              |
| Method call       | You can call methods defined in `impl` blocks on values of that type (`coin.value_in_cents()`) |

---

**In short:**  
- `impl Coin` lets you define methods for the `Coin` enum.
- `value_in_cents` is one such method, called with `coin.value_in_cents()`.
- `&self` refers to the current enum value.
- `match self` checks which variant it is (Penny, Nickel, etc.) and returns the appropriate value.




