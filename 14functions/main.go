package main

import "fmt"

func addInteger(x int, y int) int {
	return x + y
}

func proAdder(values ...int) (int, string) {
	total := 0
	for _, val := range values {
		total = total + val
	}

	return total, "Add Multiple values"
}

func main() {
	fmt.Println("This is function tutorial")
	a := addInteger(3, 4)
	fmt.Println(a)
	b, c := proAdder(3, 4, 5, 6, 8, 9)
	fmt.Println(c)
	fmt.Println(b)
}
