package main

import "fmt"

func main() {
	fmt.Println("Starting the profit calculator program")

	var revenue float64
	var taxRate float64
	var expenses float64

	fmt.Print("Enter revenue amount: ")
	fmt.Scan(&revenue)
	fmt.Print("Enter tax rate: ")
	fmt.Scan(&taxRate)
	fmt.Print("Enter expenses amount: ")
	fmt.Scan(&expenses)

	beforeTaxEarnings := revenue - expenses
	profit := beforeTaxEarnings - ((beforeTaxEarnings * taxRate) / 100)
	ratio := beforeTaxEarnings / profit

	fmt.Println("Earnings before tax: ", beforeTaxEarnings)
	fmt.Println("Profit: ", profit)
	fmt.Println("Ratio: ", ratio)
}
