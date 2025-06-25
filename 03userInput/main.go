package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to the another lecture to learn user input"
	fmt.Println(welcome)

	userInput := bufio.NewReader(os.Stdin)
	fmt.Print("Enter you name : ")

	// comma ok syntax / error ok syntax
	//   try      ,   catch
	// inputwitherr, err := userInput.ReadString('\n')
	input, _ := userInput.ReadString('\n')
	fmt.Println("Welcome Mr. :", input)
}
