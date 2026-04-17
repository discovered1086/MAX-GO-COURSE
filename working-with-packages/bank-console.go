package main

import "fmt"

func bankConsole() {
	fmt.Println("What do you want to do today?")
	fmt.Println("1. Deposit Money")
	fmt.Println("2. Withdraw Money")
	fmt.Println("3. Check Balance")
	fmt.Println("4. Exit Application")
}

func inputChoice() int {
	var choice int
	fmt.Print("Please enter your choice: ")
	fmt.Scan(&choice)
	return choice
}
