package main

import (
	"fmt"
	"sync"
)

type post struct {
	views int
	mu    sync.Mutex
}

func (p *post) inc(wg *sync.WaitGroup) {
	defer func() {
		p.mu.Unlock()
		wg.Done()
	}()
	p.mu.Lock()
	p.views += 1 //Here the modification is going on
	//p.mu.Unlock() //this should be in defer
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
 go run testing/main.go 
100
 go run testing/main.go 
100
 go run testing/main.go 
100
 go run testing/main.go 
100
*/
