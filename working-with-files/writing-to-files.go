package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const filename = "balances.txt"

func writeBalanceToFile(balance float64) {
	balanceString := fmt.Sprintf("%.2f", balance) //Convert balance to a string

	//Permission 0644 means owner of the file can read from or write to the file
	os.WriteFile(filename, []byte(balanceString), 0644)
}

func readBalanceFromFile() (float64, error) {
	bytes, err := os.ReadFile(filename)

	if err != nil {
		return 0.0, errors.New("failed to read balance from a file")
	}

	balanceString := string(bytes)
	balance, err := strconv.ParseFloat(balanceString, 64)

	if err != nil {
		return 0.0, errors.New("failed to parse balance from a file")
	}
	return balance, nil
}

func main() {
	accountBalance, err := readBalanceFromFile()

	fmt.Println("Welcome to Dummy Banking")

	if err != nil {
		fmt.Println("ERROR: Banking is not available now. Please try again")
		//panic("Banking is not available now.")
	}

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
			writeBalanceToFile(accountBalance)
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
