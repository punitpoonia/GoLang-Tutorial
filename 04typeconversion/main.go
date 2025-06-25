package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("In this we are doing type conversion")

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Please Rate our Pizza from 1 - 5 : ")

	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks for giving us rating : ", input)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	} else {

		fmt.Println("Type Conversion of Input String from user to float add 1 : ", numRating+1)
	}
}
