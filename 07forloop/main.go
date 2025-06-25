package main

import "fmt"

func main() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i++
	}

	// INFINITE LOOP

	// for {
	// 	println("hello")
	// }

	// Standard for loop

	for j := 0; j <= 5; j++ {
		fmt.Println(i)
	}

	for k := range 11 {
		fmt.Println(k)
	}
}
