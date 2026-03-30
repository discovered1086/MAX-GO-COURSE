package main

import (
	"fmt"
)

func main() {
	fmt.Println("Starting the profit calculator program")

	var revenue float64
	var taxRate float64
	var expenses float64

	//revenue, _ = strconv.ParseFloat(inputOutput("Enter revenue amount: "), 64)
	//taxRate, _ = strconv.ParseFloat(inputOutput("Enter tax rate: "), 64)
	//expenses, _ = strconv.ParseFloat(inputOutput("Enter expenses amount: "), 64)

	revenue = getUserInput("Enter revenue amount: ")
	taxRate = getUserInput("Enter tax rate: ")
	expenses = getUserInput("Enter expenses amount: ")

	fmt.Println("Revenue amount:", revenue)
	fmt.Println("Tax rate:", taxRate)
	fmt.Println("Expenses:", expenses)

	beforeTaxEarnings, profit, ratio := calculateValues(revenue, expenses, taxRate)

	fmt.Printf("Earnings before tax: %.2f \n", beforeTaxEarnings)
	fmt.Printf("Profit: %.2f \n", profit)
	fmt.Printf("Ratio: %.2f \n", ratio)
}

func getUserInput(printStatement string) float64 {
	var input float64
	fmt.Print(printStatement)
	fmt.Scan(&input)
	return input
}

func calculateValues(revenue, expenses, taxRate float64) (float64, float64, float64) {
	beforeTaxEarnings := revenue - expenses
	profit := beforeTaxEarnings - ((beforeTaxEarnings * taxRate) / 100)
	ratio := beforeTaxEarnings / profit
	return beforeTaxEarnings, profit, ratio
}
