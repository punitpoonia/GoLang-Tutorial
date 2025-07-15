package main

import (
	"fmt"
	"os"
)

func main() {
	revenue, err := getUserInput("Revenue: ")
	if err != nil {
		fmt.Println("Sorry there is some issue with your input")
		fmt.Println(err)
		panic("Try again with positive values")
	}
	expenses, err := getUserInput("Expenses: ")
	if err != nil {
		fmt.Println("Sorry there is some issue with your input")
		fmt.Println(err)
		panic("Try again with positive values")
	}
	taxRate, err := getUserInput("Tax Rate: ")
	if err != nil {
		fmt.Println("Sorry there is some issue with your input")
		fmt.Println(err)
		panic("Try again with positive values")
	}

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	fmt.Printf("%.1f\n", ebt)
	fmt.Printf("%.1f\n", profit)
	fmt.Printf("%.3f\n", ratio)
	storeResultIntoTheFile(ebt, profit, ratio)
}

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}

func storeResultIntoTheFile(rev, exp, tax float64) {
	result := fmt.Sprintf("EBT = %.1f \n Profit = %.1f \n Ratio = %.1f ", rev, exp, tax)

	os.WriteFile("Result.txt", []byte(result), 0644)

}
