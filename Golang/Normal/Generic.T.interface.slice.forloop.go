package main

import "fmt"

func printSlice[T interface{}](items []T) {
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
