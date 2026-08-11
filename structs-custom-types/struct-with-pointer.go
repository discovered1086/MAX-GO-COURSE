package main

import (
	"fmt"
	"time"
)

type userDetails struct {
	firstName string
	lastName  string
	email     string
	createdAt time.Time
}

func main() {
	inputFirstName, inputLastName, inputEmail := inputChoiceUserDetails()
	//
	//fmt.Printf("Hello %s, %s!\n", inputFirstName, inputLastName)
	//fmt.Printf("Thank you for adding %s as your email %s.\n", inputEmail, inputFirstName)

	var userInstance userDetails

	userInstance = userDetails{
		firstName: inputFirstName,
		lastName:  inputLastName,
		email:     inputEmail,
		createdAt: time.Now(),
	}

	printUserDetails(&userInstance)
}

func printUserDetails(userVar *userDetails) {
	fmt.Printf("Hello %s, %s!\n", (*userVar).firstName, userVar.lastName)
	fmt.Printf("Thank you for adding %s as your email %s.\n", userVar.email, userVar.firstName)
	fmt.Printf("Your id was created at %s.\n", userVar.createdAt.Format(time.RFC3339))
}

func inputChoiceUserDetails() (string, string, string) {
	var firstName, lastName, email string
	fmt.Print("Please enter your first name: ")
	fmt.Scan(&firstName)
	fmt.Print("Please enter your Last name: ")
	fmt.Scan(&lastName)
	fmt.Print("Please enter your email address: ")
	fmt.Scan(&email)
	return firstName, lastName, email
}
