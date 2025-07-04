package main

import "fmt"

// Generics

// we use duplicate functions for different datatype to reduce this we can use generics
// To allow only particular datatype through generic we can use [T int | string] this syntax

// WE can use generics in struct also by using the same approach

// we can use comparable instead of any or interface in generics [T comaparabel]

// WE can pass multiple types in generics also [T comparable, V string ] like this

type Stack[T any] struct { // This allow any data type in struct
	elements []T
}

// func sliceOfAnyType[T any](num []T) { // this is generic using [T any ] this syntax OR  [T interface{}]
// 	for _, val := range num {
// 		fmt.Println(val)
// 	}
// }

// func sliceOfStringANDInt[T int | string](num []T) {
// 	for _, val := range num {
// 		fmt.Println(val)
// 	}
// }

func main() {
	myStack := Stack[string]{
		elements: []string{"1", "2", "3", "4"},
	}

	fmt.Println(myStack)
	// nums := []int{1, 2, 3, 4}
	// names := []string{"punit", "poonia"}
	// sliceOfAnyType(names)
	// sliceOfStringANDInt(nums)
}
