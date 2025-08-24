keep returning a reference &y from get_val(), but as it stands, y is a local variable inside the function — which means it gets dropped as soon as the function returns, making the reference invalid.

How to fix while keeping a reference:

The reference you return must point to data that lives long enough (beyond the function call). Here are some ways to do that:

Option 1: Use a static variable

A static variable lives for the entire program lifetime ('static lifetime). You can reference it safely.

fn get_val() -> &'static i32 {
    static Y: i32 = 5;
    &Y
}

fn main() {
    let x = 10;
    let y = get_val();
    println!("x + y = {}", x + y);
}


This works because Y lives for the entire runtime of the program.

Option 2: Pass the reference into the function

If you want to refer to something outside the function, pass it as an argument:

fn get_val(y: &i32) -> &i32 {
    y
}

fn main() {
    let x = 10;
    let y_value = 5;
    let y = get_val(&y_value);
    println!("x + y = {}", x + y);
}


Here, get_val just returns the reference it received, which is valid because y_value lives in main.

Why your original code can’t work:
fn get_val() -> &i32 {
    let y = 5; // y is local and dropped after function returns
    &y        // returning reference to local variable — invalid!
}


You can’t return a reference to a local variable because Rust’s safety rules prevent dangling pointers.
