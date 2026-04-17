package main

import (
	"fmt"
	"working-with-packages/file-operations"

	"github.com/Pallinder/go-randomdata"
)

const filename = "balances.txt"

func main() {
	accountBalance, err := file_operations.ReadFloatFromFile(filename)

	fmt.Println("Welcome to Dummy Banking - " + randomdata.FullName(2))

	if err != nil {
		fmt.Println("ERROR: Banking is not available now. Please try again")
		//panic("Banking is not available now.")
	}

	for {
		bankConsole()

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
			file_operations.WriteFloatToFile(accountBalance, filename)
			return
		default:
			fmt.Println("Invalid choice, try again")
			choice = inputChoice()
		}
	}
}
