package main

import "fmt"
//any is not good use interface
func printSlice[T any](items []T) {
	for _, item := range items {
		fmt.Println(item)
	}
}

func main() {
	nums := []string{"go", "ajax", "java"}
	printSlice(nums)
	num1 := []int{1, 2, 3}
	printSlice(num1)
}
