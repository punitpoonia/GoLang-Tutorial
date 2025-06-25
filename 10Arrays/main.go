package main

import "fmt"

func main() {

	var nums [4]int

	// to check length  len(arr)
	fmt.Println(len(nums))
	fmt.Println(nums)

	// assign values

	arr := [4]int{1, 2, 3, 4}

	fmt.Println(arr)

	// 2d arrays

	arr2D := [2][2]int{{1, 2}, {3, 4}}
	fmt.Println(arr2D)

}
