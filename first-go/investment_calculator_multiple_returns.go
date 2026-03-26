package main

import (
	"fmt"
	"math"
)

const inflationRate float64 = 2.5

func main() {
	var investment, years, expectedReturnRate float64

	fmt.Print("Enter investment amount: ")
	fmt.Scan(&investment)
	fmt.Print("Enter years: ")
	fmt.Scan(&years)
	fmt.Print("Enter expected return: ")
	fmt.Scan(&expectedReturnRate)

	futureValue, futureActualReturnValue := calculateValue(investment, years, expectedReturnRate)

	fmt.Println(formatNumber(futureValue, 2, "Future value: ", "f"))
	fmt.Println(formatNumber(futureActualReturnValue, 2,
		"Future adjustment float value: ", "f"))
}

func calculateValue(investment, expectedReturnRate, years float64) (float64, float64) {
	futureValue := investment * math.Pow(1+expectedReturnRate/100, years)
	futureActualReturnValue := futureValue / math.Pow(1+inflationRate/100, years)
	return futureValue, futureActualReturnValue
}
