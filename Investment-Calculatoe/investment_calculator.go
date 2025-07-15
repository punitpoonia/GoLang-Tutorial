package main

import (
	"fmt"
	"math"
)

const INFLATIONRATE = 3.5

func calculateInvestedAmount(investedAmount, expectedReturnRate, noOfYear float64) (float64, float64) {
	fv := float64(investedAmount) * math.Pow((1+expectedReturnRate/100), float64(noOfYear))
	rfv := fv / math.Pow(1+INFLATIONRATE/100, noOfYear)
	return fv, rfv
}

func main() {

	var investedAmount float64
	expectedReturnRate := 8.5
	noOfYears := 10.0
	fmt.Print("Enter amount to be invested : ")
	fmt.Scanln(&investedAmount)
	fmt.Print("Enter expected rate of interest : ")
	fmt.Scanln(&expectedReturnRate)
	fmt.Print("Enter no of years to invest : ")
	fmt.Scanln(&noOfYears)
	futureValue, futureRealValue := calculateInvestedAmount(investedAmount, expectedReturnRate, noOfYears)
	fmt.Printf("Future Value without inflation Rs%.1f\n", futureValue)
	fmt.Printf("Future Value with inflation Rs%.1f", futureRealValue)
}
