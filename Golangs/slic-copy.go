package main
import "fmt"
func main(){
    var num = make([]int,0,5) //len is 0
    num = append(num,15)  //len of num become 1
    var num2 = make([]int,len(num))  //len of num2 become 1
     
     copy(num2,num) //copy element of num to num2
     fmt.Println(num,num2)
}


//[15] [15]
