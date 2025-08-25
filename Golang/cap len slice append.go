package main
import "fmt"

func main() {
    nums := make([]int, 2, 5) // length = 2, capacity = 5

    fmt.Println("len:", len(nums)) // 2
    fmt.Println("cap:", cap(nums)) // 5

    nums = append(nums, 200)
    
    fmt.Println("After append:")
    fmt.Println("len:", len(nums)) // 4
    fmt.Println("cap:", cap(nums)) // 5 (still within capacity)
    fmt.Println("Nums:", nums) 
    
}

/*
len: 2
cap: 5
After append:
len: 3
cap: 5
Nums: [0 0 200]

*/
