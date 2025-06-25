package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Handling Times in this tutorial : ")
	presentTime :=time.Now()
	fmt.Println(presentTime)
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday")) // By Default Values for time formatting
}
