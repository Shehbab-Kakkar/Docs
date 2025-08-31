package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("main execution started")
	fmt.Println("No. of CPUs", runtime.NumCPU())
	fmt.Println("No. of Goroutines:", runtime.NumGoroutine())
	fmt.Println("OS:", runtime.GOOS)
	fmt.Println("Arch:", runtime.GOARCH)
}
/*
main execution started
No. of CPUs 2
No. of Goroutines: 1
OS: linux
Arch: amd64
*/
