package main

import "fmt"

// LIFO
type stack[T any] struct {
	elements []T
}

func main() {
	s := stack[string]{
		elements: []string{"Go", "Rust", "Python", "Typescript"},
	}
	fmt.Println(s)

	i := stack[int]{
		elements: []int{1, 2, 4},
	}
	fmt.Println(i)
}
/* 
{[Go Rust Python Typescript]}
{[1 2 4]}
*/
