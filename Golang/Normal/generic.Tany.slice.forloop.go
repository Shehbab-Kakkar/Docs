package main

import "fmt"
//Using Type T as any or interface{}  it satisfy both int and string types
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
