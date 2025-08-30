package main

import "fmt"

func task(id int) {
	fmt.Println("doing task", id)
}

func main() {
	for i := 0; i <= 10; i++ {
		task(i)
	}
}
/*
doing task 0
doing task 1
doing task 2
doing task 3
doing task 4
doing task 5
doing task 6
doing task 7
doing task 8
doing task 9
doing task 10
*/
