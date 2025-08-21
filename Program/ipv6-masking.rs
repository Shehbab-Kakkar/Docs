fn main() {
    // Original input string
    let input_str = "fe80::f0d0:d896:3214:e4d9%39";

    // Create a mutable String to store the masked result
    let mut masked_str = String::new();

    // Iterate through each character
    for ch in input_str.chars() {
        // Check if character is a letter (a-z or A-Z)
        if ('a'..='z').contains(&ch) || ('A'..='Z').contains(&ch) {
            masked_str.push('X'); // Replace letters with 'X'
        } else {
            masked_str.push(ch);  // Keep other characters unchanged
        }
    }

    // Print the final masked string
    println!("Masked: {}", masked_str);
}
