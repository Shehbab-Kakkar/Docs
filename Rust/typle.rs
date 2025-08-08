fn main() {
    // Basic tuple example
    let tup: (i32, f64, char) = (500, 6.4, 'a');
    println!("The whole tuple is: {:?}", tup);

    // Accessing tuple elements by index
    let first = tup.0;
    let second = tup.1;
    let third = tup.2;
    println!("The first element is: {}", first);
    println!("The second element is: {}", second);
    println!("The third element is: {}", third);

    // Tuple destructuring
    let (x, y, z) = tup;
    println!("Destructured: x = {}, y = {}, z = {}", x, y, z);

    // Single-element tuple vs not-a-tuple
    let single = (5,); // tuple
    let not_a_tuple = (5); // just an integer
    println!("Single-element tuple: {:?}", single);
    println!("Not a tuple, just an integer: {}", not_a_tuple);
}

/*
The whole tuple is: (500, 6.4, 'a')
The first element is: 500
The second element is: 6.4
The third element is: a
Destructured: x = 500, y = 6.4, z = a
Single-element tuple: (5,)
Not a tuple, just an integer: 5
*/
