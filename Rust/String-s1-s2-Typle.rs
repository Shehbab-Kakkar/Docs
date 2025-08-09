fn main() {
    // Create a new String with the value "Hello"
    let s1: String = String::from("Hello");

    // Call the `callen` function, passing ownership of `s1`
    // The function returns a tuple: (original String, its length)
    let (x, y) = callen(s1);

    // Print the string and its length
    println!("{x} {y}");
}

// Define a function that takes ownership of a String
// and returns a tuple containing the same String and its length
fn callen(item: String) -> (String, usize) {
    // Get the length of the string
    let len = item.len();

    // Return a tuple with the original string and its length
    (item, len)
}
