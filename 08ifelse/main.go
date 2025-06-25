package main

import "fmt"

func main() {
	age := 16

	if age >= 18 {
		fmt.Println("Eligible for DL")
	} else {
		fmt.Println("Below 18")
	}

	// we can decalre variable inside if else loop

	if role := "admin2"; role  == "admin" {
		fmt.Println("Access Granted")
	}
}
