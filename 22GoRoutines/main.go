package main

import (
	"fmt"
	"time"
)

// To run something together or parrallely we use goroutines
// we just add go in front of fucn call

func task(id int) {
	fmt.Println("Assigned work to worker", id)
}

func main() {

	// this will not print anything because main func do not have enough time to wait for task fucntion and main function
	// executed so we have to stop main fucntion using sleep() fucntion

	for i := 0; i < 10; i++ {
		go task(i) // this is a go routines this will run parrallely this fucntion because there is no blocking
	}

	time.Sleep(time.Second * 1) // this will stop main fucntion for 1 sec to stop execution till then our go
	//  go routine will run
}
