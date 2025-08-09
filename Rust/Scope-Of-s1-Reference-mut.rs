fn main() {
    let mut s1: String = String::from("HELLO");

    {
        let w1 = &mut s1;
        w1.push_str(" WORLD");
        println!("{w1}"); // HELLO WORLD
        // w1 goes out of scope here
    }

    {
        let w2 = &mut s1;
        w2.push_str(" UNIVERSE");
        println!("{w2}"); // HELLO WORLD UNIVERSE
        // w2 goes out of scope here
    }

    // s1 is now available again if needed
    println!("{s1}"); // HELLO WORLD UNIVERSE
}

Output
HELLO WORLD
HELLO WORLD UNIVERSE
HELLO WORLD UNIVERSE
