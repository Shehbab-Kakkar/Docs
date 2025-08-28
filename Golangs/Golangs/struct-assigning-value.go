package main

import "fmt"

func main() {
	language := struct {
		new    string
		isGood bool
	}{"golang", true}

	fmt.Println(language)
}
