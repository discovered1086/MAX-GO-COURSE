package main

import (
	"fmt"
	"math"
)

func main() {
	var investment float64
	var years float64
	var expectedReturnRate float64
	const inflationRate float64 = 2.5

	fmt.Print("Enter investment amount: ")
	fmt.Scan(&investment)
	fmt.Print("Enter years: ")
	fmt.Scan(&years)
	fmt.Print("Enter expected return: ")
	fmt.Scan(&expectedReturnRate)

	futureValue := investment * math.Pow(1+expectedReturnRate/100, years)
	futureActualReturnValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println(futureValue)
	fmt.Println(futureActualReturnValue)
}
