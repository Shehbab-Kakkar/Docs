package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("testing/example.txt")
	if err != nil {
		panic(err)
	}

	fileInfo, err := f.Stat()
	if err != nil {
		panic(err)
	}
	fmt.Println("file name: ", fileInfo.Name())
	fmt.Println("file or folder: ", fileInfo.IsDir())
	fmt.Println("file size:", fileInfo.Size())
	fmt.Println("file permission:", fileInfo.Mode())
	fmt.Println("file modified at:", fileInfo.ModTime())
}

/**
file name:  example.txt
file or folder:  false
file size: 8
file permission: -rw-r--r--
file modified at: 2025-09-01 16:11:19.06916621 +0000 UTC
**/
