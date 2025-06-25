package main

import (
	"fmt"
	"slices"
)

func main() {
	// uninitilized slice by default value nil size 0
	var nums []int
	fmt.Println(nums)
	fmt.Println(len(nums))

	// Creating Slcies using make function
	var arr = make([]int, 3, 5)
	fmt.Println("initial", arr)
	fmt.Println("initial length", len(arr))    // to check the length
	fmt.Println("initial capacity ", cap(arr)) // to check maximum capacity

	// Append elements in the slices
	arr = append(arr, 1) // to add elements in the last
	arr = append(arr, 2)
	arr = append(arr, 4)
	arr = append(arr, 1)
	arr = append(arr, 2)
	arr = append(arr, 1)
	arr = append(arr, 2)
	arr = append(arr, 4)
	arr = append(arr, 4)
	arr = append(arr, 1)
	arr = append(arr, 2)
	arr = append(arr, 4)
	fmt.Println("updated", arr)
	fmt.Println("updated length", len(arr))
	fmt.Println("updated capacity", cap(arr))

	// Print within a range

	fmt.Println(arr[4:])

	// Slices.Equals(num1,num2) to compare two siles return boolean value always

	num1 := make([]int, 3)
	num2 := make([]int, 3)
	num1[0] = 1
	num1[1] = 2
	num1[2] = 3
	num2[0] = 1
	num2[1] = 2
	num2[2] = 4

	fmt.Println(slices.Equal(num1, num2))

}
