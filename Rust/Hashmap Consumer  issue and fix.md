The problem is that in your original code, you are moving `marks` into the `for` loop with `for (user, mark) in marks`, which consumes the `marks` HashMap. This means you can't use `marks` again later (like in `marks.contains_key("Markus")`), because ownership has been transferred.

To fix this, you need to iterate over a reference to `marks` (by using `&marks`) instead of consuming it. This will allow you to still use `marks` after the loop.

Here's the corrected code:

```rust
use std::collections::HashMap;

fn main() {
    let mut marks: HashMap<&str, i32> = HashMap::new();

    // Insert some data
    marks.insert("Andreas", 11);
    marks.insert("Markus", 22);
    marks.insert("Jacob", 33);

    println!("{:?}", marks);

    // Remove "Markus"
    marks.remove("Markus");
    println!("{:?}", marks);

    // Check if "Jacob" exists and print the value
    match marks.get("Jacob") {
        Some(value) => println!("Jacob's roll number is {}", value),
        None => println!("User not found in the HashMap"),
    }

    // Print the length of the HashMap
    println!("Length of HashMap is {}", marks.len());

    // Iterate over HashMap using references to avoid consuming it
    for (user, mark) in &marks {
        println!("User name is {} and marks are {}", user, mark);
    }

    // Check if "Markus" is in the HashMap
    if marks.contains_key("Markus") {
        println!("Markus is in the HashMap.");
    } else {
        println!("Markus is not in the HashMap.");
    }
}
```

### Key Fix:

* In the `for` loop: `for (user, mark) in &marks` — this avoids moving ownership of `marks` and allows it to be used later for `contains_key` and other operations.

This should compile and run without errors now.
