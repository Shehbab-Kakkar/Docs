//Progam using varadic function give me the sum of multiple integers
package main
import "fmt"
func sum(num ...int) int {
    total := 0
    for _, num := range(num){
       total = total + num
    } 
       
    return total   
}

func main() {

   num := []int{23,45,66,77}
   result := sum(num...)
   fmt.Println(result)
}
