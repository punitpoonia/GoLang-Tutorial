package main

import (
	"fmt"
	"sync"
)

// Mutex is used to protect from race condition
// Mutes locked the resoirce so that one go routine can access the resource at one time

type post struct {
	views int
	mu    sync.Mutex
}

func (p *post) increment(wg *sync.WaitGroup) {
	defer wg.Done()
	p.mu.Lock()
	p.views = p.views + 1
	defer p.mu.Unlock()
}

func main() {
	var wg sync.WaitGroup
	myPost := post{views: 0}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go myPost.increment(&wg)

	}

	wg.Wait()
	fmt.Println(myPost.views)
}
