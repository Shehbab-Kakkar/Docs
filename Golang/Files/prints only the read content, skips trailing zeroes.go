package main

import (
	"os"
)

func main() {
	//read file
	f, err := os.Open("testing/example.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	fileInfo, err := f.Stat()
	if err != nil {
		// If an error occurs while getting file info, terminate the program
		panic(err)
	}

	buf := make([]byte, fileInfo.Size())
	d, err := f.Read(buf) // read the length of the buffer which is 10
	if err != nil {
		panic(err)
	}
	for i := 0; i < len(buf); i++ {
		println("data", d, string(buf[:i])) //prints only the read content, skips trailing zeroes
	}
}
/*
data 8 
data 8 t
data 8 te
data 8 tes
data 8 test
data 8 test 
data 8 test o
data 8 test ok
*/
