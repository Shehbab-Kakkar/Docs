//print the element of the slice 
package main
import "fmt"
func main(){
  slicePrint := []int{2,4,7}
  for i :=0;i<len(slicePrint);i++ {
    fmt.Println(slicePrint[i])
  }
  fmt.Println("...")  
  for _,r := range(slicePrint){
    println(r)
  }
}
