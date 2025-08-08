fn main() {
    // Declare an array named 'numbers' of type i32 (32-bit signed integer)
    // with a fixed size of 3, initialized with values 1, 2, and 3.
    let numbers: [i32; 3] = [1, 2, 3];

    // Iterate over each element in the 'numbers' array.
    // The '&' indicates that 'num' is a reference to an element,
    // avoiding copying the value for each iteration.
    for num in &numbers {
        // Print each number followed by a space.
        // The '{}' is a placeholder for the value of 'num', and '!s' indicates
        // that the value should be formatted for display.
        print!("{} ", num);
    }
    // Print a newline character at the end to move to the next line after printing all numbers.
    println!();
}
