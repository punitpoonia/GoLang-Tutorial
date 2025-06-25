package main

import "fmt"

func changeNum(num *int) {
	*num = 5
	fmt.Println(num)
}

func main() {

	num := 1
	changeNum(&num)
	fmt.Println("Number after change : ", num)

}
