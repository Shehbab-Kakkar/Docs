fn main(){
  let vector:Vec<i32>= vec![23,45,67];
  vector.clear();
  println!("{:?}",vector);
  println!("{}",vector.len());
}// Error
fn main(){
  let mut vector:Vec<i32>= vec![23,45,67];
  vector.clear();
  println!("{:?}",vector);
  println!("{}",vector.len());
} //
/*
[]
0
*/
