package main
import "fmt"
func main(){
    var nums = make([]int,0,5)
    fmt.Println(nums)
    fmt.Println(len(nums))
    fmt.Println(cap(nums))

}

/*
[]
0
5
*/
