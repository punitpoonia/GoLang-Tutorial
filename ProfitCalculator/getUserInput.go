package main

import (
	"errors"
	"fmt"
)

func getUserInput(infoText string) (float64, error) {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)

	if userInput <= 0 {
		return 0, errors.New("value can not be zero or negative")
	}
	return userInput, nil
}
