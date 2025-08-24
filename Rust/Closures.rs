fn main(){
    let add = |a,b| a+b;
    println!("Add: {}",add(2,3));
    
    let product  = || println!("Hello world");
    product();
    
    let product1 = |a,b| {
        let sum = a+b;
        let mult = a*b;
        sum + mult
    };
    println!("{}",product1(3,4));
}
