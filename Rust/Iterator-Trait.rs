// Counter struct with a single field of type u32
struct Counter {
    count: u32,
}

impl Counter {
    // Constructor method that initializes Counter with count = 0
    fn new() -> Counter {
        Counter { count: 0 }
    }
}

// Implement the Iterator trait for Counter
impl Iterator for Counter {
    // Set the associated Item type to u32
    type Item = u32;
    
    // Implement the next() method
    fn next(&mut self) -> Option<Self::Item> {
        // Increment count each time it is called
        self.count += 1;
        
        // Return Some(count) while count <= 5
        if self.count <= 5 {
            Some(self.count)
        } else {
            // Return None once the counter exceeds 5
            None
        }
    }
}

fn main() {
    // Create a new counter instance
    let counter = Counter::new();
    
    // Use the iterator to print numbers 1 through 5 to the console
    println!("Printing numbers 1 through 5:");
    
    for number in counter {
        println!("{}", number);
    }
    
    // Alternative way using iterator methods
    println!("\nUsing collect() to gather all values:");
    let counter2 = Counter::new();
    let numbers: Vec<u32> = counter2.collect();
    println!("{:?}", numbers);
    
    // Another alternative using while let
    println!("\nUsing while let pattern:");
    let mut counter3 = Counter::new();
    while let Some(number) = counter3.next() {
        println!("{}", number);
    }
}