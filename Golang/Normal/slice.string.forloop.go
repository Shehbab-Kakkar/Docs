package main

import "fmt"

func printSlice(items []string) {
	for _, item := range items {
		fmt.Println(item)
	}
}

func main() {
	nums := []string{"go", "ajax", "java"}
	printSlice(nums)
}
