package main

import "fmt"

func main() {
	m := map[string]int{"Ram": 40, "Sham": 23}
	_, ok := m["Ramq"]
	if ok {
		fmt.Println("All ok")
	} else{
	  fmt.Println("not ok")
	}
}
