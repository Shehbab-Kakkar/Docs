fn main(){
    let s1: String = String::from("Hello");
    let (x,y) = cal_Stringlen(s1);
    println!("Print s1 is {}, len of s1 {}",x,y);
}


fn cal_Stringlen(item: String) -> (String, u8){
    let result = item.len() as u8;
    (item, result)
}


Standard Output
Print s1 is Hello, len of s1 5
