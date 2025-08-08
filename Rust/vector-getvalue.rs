fn main() {
    let vector_name = vec![10, 20, 30];
    println!("{:?}", vector_name);

    let mut another_vector = vec![30, 10];
    println!("{:?}", another_vector);

    if let Some(value) = vector_name.get(2) {
        println!("Value: {}", value);
    }
}
