fn main() {
    // Step 1: Create the original string
    let s1 = String::from("Hello");

    // Step 2: Pass the original string by immutable reference
    //         to a function that will return a new modified version
    let changed = cal_len(&s1);

    // Step 3: Print both the original and the modified string
    println!("Original = {}, Changed = {}", s1, changed);
}

// This function takes an immutable reference to a String
// and returns a new String that appends "World" to it
fn cal_len(item: &String) -> String {
    // Step 4: Create a new String (empty for now)
    let mut modified = String::new();

    // Step 5: Append the original string's contents
    modified.push_str(item); // "Hello"

    // Step 6: Append "World" to it
    modified.push_str("World"); // "HelloWorld"

    // Step 7: Return the new String
    modified
}


Original = Hello, Changed = HelloWorld
