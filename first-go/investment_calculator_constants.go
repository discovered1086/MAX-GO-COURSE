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

	//fmt.Println("Future value:", futureValue)
	//fmt.Println("Future adjustment value:", futureActualReturnValue)

	//Formatted numbers
	//fmt.Printf("Future value: %.2f\n", futureValue)
	//fmt.Printf("Future adjustment float value: %f\n", futureActualReturnValue)
	//fmt.Printf("Future adjustment value: %v\n", futureActualReturnValue)
	//fmt.Printf("Future value type: %T\n", futureValue)
	//fmt.Printf("Future value with zero decimal places: %.0f\n", futureValue)

	formattedFutureValue := fmt.Sprintf("%.2f", futureValue)
	formattedFutureActualReturnValue := fmt.Sprintf("Future adjustment float value: %.2f\n ", futureActualReturnValue)
	futureValueType := fmt.Sprintf("%T", futureValue)
	formattedWithZeroDecimalValue := fmt.Sprintf("%.f", futureValue)
	formattedNewLineValue := fmt.Sprintln(futureValue)

	fmt.Println("Future value:", formattedFutureValue)
	fmt.Print(formattedFutureActualReturnValue)
	fmt.Println("Future adjustment value: ", futureActualReturnValue)
	fmt.Println("Future value type:", futureValueType)
	fmt.Println("Future value with zero decimal places: ", formattedWithZeroDecimalValue)
	fmt.Print("Future value in new line:", formattedNewLineValue)

	fmt.Printf(`Future value: %.2f
		 and future adjustment value in new lines %.2f`, futureValue, futureActualReturnValue)
}
