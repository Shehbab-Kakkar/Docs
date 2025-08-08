//Find the Player Direction
enum Direction{
        up,
        down,
        right,
        left,
    }
fn main(){
   let player_direction = Direction::up;
   match player_direction {
       Direction::up => println!("Direction of Player is up"),
       Direction::down => println!("Direction of Player is down"),
       Direction::right => println!("Direction of Player is right"),
       Direction::left => println!("Direction of Player is left"),
   }

}
