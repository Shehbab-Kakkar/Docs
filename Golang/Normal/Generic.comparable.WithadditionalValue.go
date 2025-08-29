package main

import "fmt"

func printSlice[T comparable, V string](items []T, name V) { // string is the addition value.
	for _, item := range items {
		fmt.Println(item, name)
	}
}

func main() {
	// s := []string{"go", "ajax", "java"}
	// printSlice(s)
	// i := []int{1, 2, 3}
	// printSlice(i)
	b := []bool{true, false}
	printSlice(b, "ram")

}
