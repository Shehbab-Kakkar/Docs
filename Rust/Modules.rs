// Main module
mod movie {
    // Nested module within movie
    pub mod film {
        pub fn films1() {
            println!("Film 1: The Shawshank Redemption");
        }
    }
    
    pub fn films_under_movie() {
        println!("Listing films under movie module:");
        film::films1();
    }
}

mod nation {
    // Nested module within nation
    pub mod film {
        pub fn films1() {
            println!("Film 1: The Godfather");
        }
    }
    
    pub fn films_under_nation() {
        println!("Listing films under nation module:");
        film::films1();
    }
}

// Main function
fn main() {
    println!("Welcome to the Film Database!\n");
    
    movie::films_under_movie();
    nation::films_under_nation();
}
