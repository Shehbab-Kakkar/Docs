package main
import "fmt"
func main() {
  m := map[string]int{"Ram":40,"Sita":23}
  for k,v := range(m) {
    fmt.Println(k,v)
  }
}
