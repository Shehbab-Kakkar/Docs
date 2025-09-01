package main

import (
	"fmt"
	"sync"
)

type post struct {
	views int
}

func (p *post) inc(wg *sync.WaitGroup) {
	defer wg.Done()
	p.views += 1    //Single resource is updated by multiple process
}

func main() {
	var wg sync.WaitGroup
	myPost := post{views: 0}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go myPost.inc(&wg)
	}
	wg.Wait()
	fmt.Println(myPost.views)
}
/*
$ go run testing/main.go 
100
 ~]$ go run testing/main.go 
95
$ go run testing/main.go 
100

]$ go run testing/main.go 
99

[sk-dev@sk
*/
