package main
import "fmt"
func main(){
  m := []int{34,55,66}
  for i,j := range(m){
    fmt.Println(i,j)
  }
}
