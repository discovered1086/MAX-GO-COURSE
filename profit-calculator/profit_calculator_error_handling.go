package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

// Goals
//  1. Validate user input
//     ==> Show error message and exit if invalid input is provided
//     No negative numbers
//     Not 0
//  2. Store calculated values in a file

func writeValueToFile(beforeTaxEarnings, profit, ratio float64) {
	fileName := "profit_calculator_error_handling.txt"

	fileValue := "Before tax earnings: " +
		strconv.FormatFloat(beforeTaxEarnings, 'f', 2, 64) +
		"\nProfit: " + strconv.FormatFloat(profit, 'f', 2, 64) +
		"\nRatio: " + strconv.FormatFloat(ratio, 'f', 2, 64)

	os.WriteFile(fileName, []byte(fileValue), 0644)
}
func main() {
	fmt.Println("Starting the profit calculator program")

	revenue, err := getInput("Enter revenue amount: ")
	taxRate, err := getInput("Enter tax rate: ")
	expenses, err := getInput("Enter expenses amount: ")

	if err != nil {
		fmt.Println(err)
		panic("Exiting the application due to an error")
	}

	fmt.Println("Revenue amount:", revenue)
	fmt.Println("Tax rate:", taxRate)
	fmt.Println("Expenses:", expenses)

	beforeTaxEarnings, profit, ratio := calculate(revenue, expenses, taxRate)

	fmt.Printf("Earnings before tax: %.2f \n", beforeTaxEarnings)
	fmt.Printf("Profit: %.2f \n", profit)
	fmt.Printf("Ratio: %.2f \n", ratio)

	writeValueToFile(beforeTaxEarnings, profit, ratio)
}

func getInput(printStatement string) (float64, error) {
	var input float64
	fmt.Print(printStatement)
	fmt.Scan(&input)

	if input <= 0 {
		return input, errors.New("ERROR: invalid input...!!")
	}

	return input, nil
}

func calculate(revenue, expenses, taxRate float64) (float64, float64, float64) {
	beforeTaxEarnings := revenue - expenses
	profit := beforeTaxEarnings - ((beforeTaxEarnings * taxRate) / 100)
	ratio := beforeTaxEarnings / profit
	return beforeTaxEarnings, profit, ratio
}
