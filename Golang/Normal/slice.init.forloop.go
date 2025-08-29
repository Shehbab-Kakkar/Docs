package main

import "fmt"

func printSlice(items []int) {
	for _, item := range items {
		fmt.Println(item)
	}
}

func main() {
	nums := []int{2, 3, 4}
	printSlice(nums)
}
