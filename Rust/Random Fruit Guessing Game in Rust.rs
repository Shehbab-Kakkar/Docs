/*
===========================================
🍌 Random Fruit Guessing Game in Rust 🕹️
===========================================

▶ Description:
This is a simple terminal-based guessing game written in Rust.

The program randomly selects one fruit from a predefined list: 
["graph", "banana", "orange"].

The user is prompted to guess the fruit. The program:
- Accepts input from the user
- Checks if the input is a valid fruit from the list
- Compares the guess to the randomly selected fruit
- Provides feedback:
    - If the guess is correct, it congratulates the user and ends the game
    - If incorrect, it allows the user to try again

▶ Features:
- Uses random number generation (`rand` crate)
- Reads input from the terminal (`std::io`)
- Compares strings in a case-insensitive way
- Loops until the correct fruit is guessed
- Handles invalid inputs gracefully

▶ Usage:
- Run the program
- Type a fruit name (e.g., banana) and press Enter
- Keep trying until you guess the correct fruit

Have fun guessing! 🎯

*/



// Import the standard I/O library for reading user input
use std::io;

// Import random number generation functionality
use rand::Rng;
use rand::thread_rng;

// Function to check if the guessed fruit is the same as the randomly selected one
fn guess_checker(guess: &str, actual: &str) -> bool {
    guess == actual
}

fn main() {
    // Define a list of fruits to choose from
    let guess_list = ["graph", "banana", "orange"];

    // Initialize a random number generator
    let mut rng = thread_rng();

    // Randomly pick an index from the fruit list
    let index = rng.gen_range(0..guess_list.len());

    // Select the fruit at the random index
    let random_fruit = guess_list[index];

    // Start an infinite loop to keep asking the user for input
    loop {
        // Create a new mutable string to store user input
        let mut input = String::new();

        // Ask the user to guess the fruit
        println!("Guess the fruit (options: graph, banana, orange):");

        // Read the user's input from the terminal
        match io::stdin().read_line(&mut input) {
            Ok(_) => {
                // Clean the input: remove newline and convert to lowercase
                let fruit_selected = input.trim().to_lowercase();

                // Show the user what they entered
                println!("Fruit Selected: {}", fruit_selected);

                // Check if the entered fruit is in the guess list
                if !guess_list.contains(&fruit_selected.as_str()) {
                    println!("Fruit entered is not in the list. Try again with one of the listed fruits.");
                    continue; // Ask again
                }

                // Use the guess_checker function to compare guess and actual fruit
                if guess_checker(&fruit_selected, random_fruit) {
                    println!("🎉 You are the winner!");
                    break; // Exit the loop if correct
                } else {
                    println!("❌ Incorrect. Try again.");
                }
            }

            // Handle any error that occurs while reading input
            Err(error) => {
                println!("Error reading input: {}", error);
            }
        }
    }
}
