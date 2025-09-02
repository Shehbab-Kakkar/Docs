package main

import (
	"os"
)

func main() {
	f, err := os.Create("testing/example3.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	f.WriteString("First Line\n")
	f.WriteString("Second Line\n")
	byte := []byte("Hello Golang")
	f.Write(byte)
}
