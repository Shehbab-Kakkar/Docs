trait HasArea {
    fn area(&self) -> i64;
}

struct Square {
    side: i64,
}

impl HasArea for Square {
    fn area(&self) -> i64 {
        self.side * self.side
    }
}

struct Circle {
    radius: i64,
}

impl HasArea for Circle {
    fn area(&self) -> i64 {
        self.radius * self.radius * 3 // Approximation of Pi = 3 for simplicity
    }
}

fn print_area<T: HasArea>(shape: T) {
    println!("This shape has an area of {}", shape.area());
}

fn main() {
    let square = Square { side: 4 };
    let circle = Circle { radius: 3 };

    print_area(square);  // Here `T` is `Square`
    print_area(circle);  // Here `T` is `Circle`
}
