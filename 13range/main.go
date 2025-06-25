package main

import "fmt"

// iterating over data structures range is used

func main() {
	sum := 0
	nums := []int{1, 2, 3, 4}
	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
		sum += nums[i]
	}

	fmt.Println("Sum of all values is : ", sum)

	// Using Range keyn word for slices

	for i, num := range nums {
		fmt.Println(i, num)

	}

	// Range for maps
	// by deault if only one variable in for range it will take key name

	m := map[string]string{"name": "punit", "gender": "male"}

	for k, v := range m {
		fmt.Println(k, v)
	}

	// Range for strings
	// unicodes value is given here or rune

	for j, g := range "golang" {
		fmt.Println(j, string(g))
	}
}
