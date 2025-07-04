package main

import "fmt"

func main() {
	channel1 := make(chan string)
	channel2 := make(chan int)

	go func() {
		channel2 <- 10
	}()

	go func() {
		channel1 <- "punitpoonia"
	}()

	for i := 0; i < 2; i++ {
		select {
		case channel1val := <-channel1:
			fmt.Println("Received data from channel 1", channel1val)

		case chan2 := <-channel2:
			fmt.Println("receiving data from channel 2,", chan2)
		}
	}
}
