/*
[dependencies]

*/
use std::collections::HashMap;
fn main() {
    println!("HashMap Progam Insert and Update");
    let mut marks: HashMap<&str, i32> = HashMap::new();
    marks.insert("Ram",23);
    marks.insert("Sahil",24);
    marks.insert("Arjun",25); 
    println!("Print the marks\n{:?}",marks);
    if let Some(mark) = marks.get("Ram") {
        println!("Ram is present his marks : {}",mark);
        println!("Updating Ram value");
        marks.insert("Ram",50);
        println!("Ram value updated to {}",marks["Ram"]);
        
    } else {
        println!("Ram is not Present")
    }         
    marks.entry("Jassy").or_insert(40);
    marks.entry("Arjun").and_modify(|i| *i += 500);
    println!("{:?}",marks);
}
