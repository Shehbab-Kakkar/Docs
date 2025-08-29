package main

import "fmt"

func printSlice[T int | string | bool](items []T) {
	for _, item := range items {
		fmt.Println(item)
	}
}

func main() {
	s := []string{"go", "ajax", "java"}
	printSlice(s)
	i := []int{1, 2, 3}
	printSlice(i)
	b := []bool{true, false}
	printSlice(b)

}
