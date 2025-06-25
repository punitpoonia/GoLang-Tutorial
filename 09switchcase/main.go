package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	switch i {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("two")

	}

	// Multiple switch condition

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Its Weekend enjoy your day")
	default:
		fmt.Println("Its working day")
	}

	// Type Switch

	whoAMI := func(i any) {
		switch t := i.(type) {
		case int:
			fmt.Println("Integer")
		case bool:
			fmt.Println("Boolean")
		case string:
			fmt.Println("String")
		case float32:
			fmt.Println("Float")
		default:
			fmt.Println("other", t)
		}
	}

	whoAMI("Punit")
	whoAMI(34)

}
