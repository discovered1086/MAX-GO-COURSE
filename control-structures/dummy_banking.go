package main

import "fmt"

func main() {
	accountBalance := 1000.0

	fmt.Println("Welcome to Dummy Banking")
	fmt.Println("What do you want to do today?")
	fmt.Println("1. Deposit Money")
	fmt.Println("2. Withdraw Money")
	fmt.Println("3. Check Balance")
	fmt.Println("4. Exit Application")

	choice := inputChoice()

	userWantsBalance := choice == 3
	userWantsDeposit := choice == 1
	userWantsWithdraw := choice == 2
	userWantsToExit := choice == 4

	if userWantsBalance {
		fmt.Println("Your balance is ", accountBalance)
	} else if userWantsDeposit {
		var depositAmount float64
		fmt.Print("Please enter deposit amount: ")
		fmt.Scan(&depositAmount)

		//Validate deposit amount
		if depositAmount <= 0 {
			fmt.Println("Invalid deposit amount")
			return
		}
		accountBalance += depositAmount
		fmt.Println("Deposit successfully completed! Your balance is ", accountBalance)
	} else if userWantsWithdraw {
		var withdrawAmount float64
		fmt.Print("Please enter withdrawal amount: ")
		fmt.Scan(&withdrawAmount)

		//Check account balance
		if withdrawAmount <= 0 {
			fmt.Println("Invalid Withdrawal amount")
			return
		}

		if withdrawAmount > accountBalance {
			fmt.Println("Insufficient funds")
			return
		}
		accountBalance -= withdrawAmount
		fmt.Println("Withdrawal successfully completed! Your balance is ", accountBalance)
	} else if userWantsToExit {
		fmt.Println("Exiting Application")
	} else {
		fmt.Println("Invalid choice, try again")
		choice = inputChoice()
	}
}

func inputChoice() int {
	var choice int
	fmt.Print("Please enter your choice: ")
	fmt.Scan(&choice)
	return choice
}

/******* EXAMPLES WITH RECURSION****************
func checkBalance(accountBalance, withdrawalAmount float64) {
	if withdrawalAmount > accountBalance {
		fmt.Println("Insufficient funds, please enter a different withdrawal amount")
		fmt.Scan(&withdrawalAmount)
		checkBalance(accountBalance, withdrawalAmount)
	}
	return
}

func validateDepositAmount(depositAmount float64) {
	if depositAmount <= 0 {
		fmt.Println("Invalid deposit amount, please enter a valid deposit amount")
		fmt.Scan(&depositAmount)
		validateDepositAmount(depositAmount)
	}
	return
}
*******************************/
