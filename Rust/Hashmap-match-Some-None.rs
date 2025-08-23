use std::collections::HashMap;
fn main(){
    let mut useDetails: HashMap<&str,i32> = HashMap::new();
    useDetails.insert("John",22);
     useDetails.insert("Ronny",23);
     useDetails.insert("Mark",24);
     match useDetails.get("Ronny") {
         Some(value) => println!("Ronny user exist and his Roll No:{}",value), 
         None => println!("Ronny user do not exist"),
     }
}
