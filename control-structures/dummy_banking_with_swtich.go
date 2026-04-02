package main

import "fmt"

func main() {
	accountBalance := 1000.0

	fmt.Println("Welcome to Dummy Banking")

	for {
		fmt.Println("What do you want to do today?")
		fmt.Println("1. Deposit Money")
		fmt.Println("2. Withdraw Money")
		fmt.Println("3. Check Balance")
		fmt.Println("4. Exit Application")

		choice := inputChoice()

		switch choice {
		case 1:
			var depositAmount float64
			fmt.Print("Please enter deposit amount: ")
			fmt.Scan(&depositAmount)

			//Validate deposit amount
			if depositAmount <= 0 {
				fmt.Println("Invalid deposit amount")
				continue
			}
			accountBalance += depositAmount
			fmt.Println("Deposit successfully completed! Your balance is ", accountBalance)
		case 2:
			var withdrawAmount float64
			fmt.Print("Please enter withdrawal amount: ")
			fmt.Scan(&withdrawAmount)

			//Check account balance
			if withdrawAmount <= 0 {
				fmt.Println("Invalid Withdrawal amount")
				continue
			}

			if withdrawAmount > accountBalance {
				fmt.Println("Insufficient funds")
				continue
			}
			accountBalance -= withdrawAmount
			fmt.Println("Withdrawal successfully completed! Your balance is ", accountBalance)
		case 3:
			fmt.Println("Your balance is ", accountBalance)
		case 4:
			fmt.Println("Exiting Application")
			fmt.Println("Thank you for choosing our bank..!!")
			return
		default:
			fmt.Println("Invalid choice, try again")
			choice = inputChoice()
		}
	}
}

func inputChoice() int {
	var choice int
	fmt.Print("Please enter your choice: ")
	fmt.Scan(&choice)
	return choice
}
