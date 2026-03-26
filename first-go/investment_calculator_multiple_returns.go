package main

import (
	"fmt"
	"math"
	"strconv"
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

	futureValue, futureActualReturnValue := calculateValueDifferentReturn(investment, years, expectedReturnRate)

	fmt.Println(printNumber(futureValue, 2, "Future value: ", "f"))
	fmt.Println(printNumber(futureActualReturnValue, 2,
		"Future adjustment float value: ", "f"))
}

//func calculateValue(investment, expectedReturnRate, years float64) (float64, float64) {
//	futureValue := investment * math.Pow(1+expectedReturnRate/100, years)
//	futureActualReturnValue := futureValue / math.Pow(1+inflationRate/100, years)
//	return futureValue, futureActualReturnValue
//}

// Alternative easier approach to returns
func calculateValueDifferentReturn(investment, expectedReturnRate, years float64) (futureValue float64, futureActualReturnValue float64) {
	futureValue = investment * math.Pow(1+expectedReturnRate/100, years)
	futureActualReturnValue = futureValue / math.Pow(1+inflationRate/100, years)
	return futureValue, futureActualReturnValue
}

func printNumber(num float64, decimalPlace int, prefix, formattingChar string) string {
	formatter := "%." + strconv.Itoa(decimalPlace) + formattingChar
	return fmt.Sprintf(prefix+formatter, num)
}
