package main
import "fmt"
func main(){
    var nums = make([]int,2,5)
    fmt.Println(nums)
    nums = append(nums,10)
    nums = append(nums,20)
    nums = append(nums,30)
    
    fmt.Println(nums)  
    fmt.Println(cap(nums))
    nums = append(nums,40)
    fmt.Println(nums)  
    fmt.Println(cap(nums))
}

/*
[0 0]
[0 0 10 20 30]
5
[0 0 10 20 30 40]
10

*/
