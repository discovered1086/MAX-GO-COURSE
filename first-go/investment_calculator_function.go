package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	investment, years, expectedReturnRate := 1000.0, 10.0, 5.5
	const inflationRate float64 = 2.5

	futureValue := investment * math.Pow(1+expectedReturnRate/100, years)
	futureActualReturnValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println(formatNumber(futureValue, 2, "Future value: ", "f"))
	fmt.Println(formatNumber(futureActualReturnValue, 2,
		"Future adjustment float value: ", "f"))
}

func formatNumber(num float64, decimalPlace int, prefix, formattingChar string) string {
	formatter := "%." + strconv.Itoa(decimalPlace) + formattingChar
	return fmt.Sprintf(prefix+formatter, num)
}
