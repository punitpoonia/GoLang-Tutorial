package main

import (
	"fmt"
	"time"
)

// unbuffered channel = we can send one data at a time
// buffered channel =  we can send limited data without blocking

// channels also use for syncronization of go routines

// channel is like a pipe from one side we send the data and on the other end we receive the data
// when multiple go routines we are working needs to handle them we use channels
// to communicate between go routines
// sending data
// func reciveMsgFromChannel(num chan int) {
// 	fmt.Println("Message recives from channel ", <-num)
// }

// //recieving channel

// func sum(result chan int, num1 int, num2 int) {
// 	numresult := num1 + num2
// 	result <- numresult

// }

func emailSending(email chan string, done chan bool) {
	defer func() { done <- true }()
	for val := range email {
		fmt.Println("Sending email to : ", val)
		time.Sleep(time.Second)
	}

}

func main() {

	// buffered channel

	bufferedChannel := make(chan string, 100)
	done := make(chan bool)

	go emailSending(bufferedChannel, done)
	for i := 0; i < 5; i++ {
		bufferedChannel <- fmt.Sprintf("%d@gmail.com", i)
	}
	// bufferedChannel <- "punit@gmail.com"
	// bufferedChannel <- "ram@gmail.com"

	// fmt.Println(<-bufferedChannel)
	// fmt.Println(<-bufferedChannel)

	// result := make(chan int)
	// number := make(chan int)
	// go reciveMsgFromChannel(number)
	// go sum(result, 4, 5)

	// number <- 5

	// res := <-result
	close(bufferedChannel) // closing the channel in important
	// lets create a channel first
	// fmt.Println("sum of the numbers from channel ", res)
	<-done
	// myChannel := make(chan string)

	// myChannel <- "Ping" // data sending to channel <-

	// // channel are blocking untill the second side is ready to recive the msg

	// msg := <-myChannel // here we are reciving the data from channel <- then channel name

	// fmt.Println(msg)
}
