package main
import "fmt"
func main(){
    var num = make([]int,5,5)// If Len is 0 then we can't add
    fmt.Println(num)
    num[0]=1
    num[1]=2
    num[2]=3
    num[3]=4
    fmt.Println(num)
    fmt.Println(len(num))
    fmt.Println(cap(num))

}

[0 0 0 0 0]
[1 2 3 4 0]
5
5
