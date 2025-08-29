package main

import "fmt"


// LIFO
type stack struct {
	elements []int
}

func main() {
	myStack := stack{
		elements: []int{1, 2, 4},
	}
	fmt.Println(myStack)

}
/*{[1 2 4]}*/
