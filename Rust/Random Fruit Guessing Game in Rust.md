# 🍌 Random Fruit Guessing Game in Rust 🕹️

---

## 📝 Description

A simple terminal-based guessing game written in Rust.  
The program randomly selects one fruit from a predefined list (`["graph", "banana", "orange"]`).  
You, the user, are prompted to guess the fruit. The game continues until you guess correctly.

---

## 🚦 Program Flow

1. **Fruit List Initialization:**  
   The game starts by defining a list of fruits: `["graph", "banana", "orange"]`.

2. **Random Selection:**  
   - A random index is generated using the `rand` crate.
   - The fruit at that index becomes the "secret" fruit to guess.

3. **Game Loop:**  
   - The program repeatedly prompts the user to enter a guess.
   - User input is read, cleaned (trimmed & lowercased), and checked:
     - If the input is **not** in the fruit list, the user is told to try again.
     - If the guess **matches** the random fruit, the user wins and the game ends.
     - If the guess **does not match**, the user is encouraged to try again.
   - The loop continues until the correct fruit is guessed.

---

## 🧩 Key Libraries & Features

### 1. **Standard Library (`std::io`):**
   - Used for reading user input from the terminal.
   - `io::stdin().read_line(&mut input)` reads a line from the user.

### 2. **Random Number Generation (`rand` crate):**
   - `rand::Rng` and `rand::thread_rng` are used for generating random numbers.
   - Picks a random index from the fruit list.

### 3. **String Handling:**
   - `.trim()` removes whitespace/newlines from input.
   - `.to_lowercase()` ensures case-insensitive comparison.

### 4. **Function Usage:**
   - `guess_checker(guess: &str, actual: &str) -> bool`:  
     Compares the user's guess and the actual fruit for equality.

### 5. **Error Handling:**
   - The result of `read_line()` is matched with `Ok(_)` and `Err(error)`:
     - `Ok(_)`: The input was read successfully; proceed to process it.
     - `Err(error)`: Something went wrong while reading input; print the error message.

---

## 🔄 The Main Loop

| Step                           | What Happens                                                         |
|---------------------------------|---------------------------------------------------------------------|
| Prompt for input                | The user sees: `Guess the fruit (options: graph, banana, orange):`  |
| Read and clean user input       | Input is trimmed and lowercased for comparison                      |
| Input validation                | If not in fruit list, prompt again                                  |
| Guess comparison                | If correct, print winner message and exit                           |
| Incorrect guess                 | Print encouragement and repeat                                      |
| Input error                     | Print error and continue                                            |

---

## 🦾 Example Run

```text
Guess the fruit (options: graph, banana, orange):
> Banana
Fruit Selected: banana
❌ Incorrect. Try again.
Guess the fruit (options: graph, banana, orange):
> orange
Fruit Selected: orange
🎉 You are the winner!
```

---

## 🚩 How to Run

1. Make sure you have [Rust](https://www.rust-lang.org/tools/install) and the `rand` crate (`cargo add rand`).
2. Save the code in a `.rs` file.
3. Run with `cargo run` or `rustc yourfile.rs && ./yourfile`.
4. Type your guesses in the terminal and enjoy the game!

---

## 🛠️ Customization

- Add more fruits to the list for extra challenge.
- Make guesses case-insensitive (already handled).
- Provide a hint system or limit attempts for advanced gameplay.

---

## 💡 Learning Points

- Working with random numbers in Rust (`rand`)
- Reading and handling terminal input (`std::io`)
- Error handling using `Result`, `Ok`, and `Err`
- Function usage and code organization
- Looping and control flow
- Basic string manipulation

---

Happy guessing! 🎯
