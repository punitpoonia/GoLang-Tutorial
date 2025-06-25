package main

import "fmt"

func anyTypeFunction(nums ...int) int {
	total := 0
	for _, val := range nums {
		total = total + val
	}

	return total

}

func main() {

	nums := []int{1, 2, 3, 4, 5, 6}
	result := anyTypeFunction(nums...)
	fmt.Println(result)

}
