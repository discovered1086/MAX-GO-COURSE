package main

import (
	"fmt"
	"math"
)

func main() {
	investment, years, expectedReturnRate := 1000.0, 10.0, 5.5
	const inflationRate float64 = 2.5

	futureValue := investment * math.Pow(1+expectedReturnRate/100, years)
	futureActualReturnValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println(futureValue)
	fmt.Println(futureActualReturnValue)
}
