package main

import "fmt"
//comparable is an interface that is implemented by all comparable types 
//(booleans, numbers, strings, pointers, channels, arrays of comparable types, structs whose fields are all comparable types). 
//The comparable interface may only be used as a type parameter constraint, not as the type of a variable.
func printSlice[T comparable](items []T) {
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
/*
go
ajax
java
1
2
3
true
false
*/
