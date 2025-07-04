package main

import (
	"fmt"
	"sync"
)

// waitgroup needs to remeber three things add the waitgroup done the waitgroup and wait for the waitgroup

func task(id int, w *sync.WaitGroup) {
	defer w.Done() // defer run after execution of your function
	fmt.Println("Assigned work to worker", id)

}

func main() {
	var wg sync.WaitGroup
	for i := range 10 {
		wg.Add(1)
		go task(i, &wg)
	}
	wg.Wait()
}
