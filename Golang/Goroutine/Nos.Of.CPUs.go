package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("main execution started")
	fmt.Println("No. of CPUs", runtime.NumCPU())

}
