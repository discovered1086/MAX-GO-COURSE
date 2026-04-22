package main

import (
	"fmt"
	"time"
)

type user struct {
	firstName string
	lastName  string
	email     string
	createdAt time.Time
}

func main() {
	inputFirstName, inputLastName, inputEmail := inputChoice()
	//
	//fmt.Printf("Hello %s, %s!\n", inputFirstName, inputLastName)
	//fmt.Printf("Thank you for adding %s as your email %s.\n", inputEmail, inputFirstName)

	var userInstance user

	userInstance = user{
		firstName: inputFirstName,
		lastName:  inputLastName,
		email:     inputEmail,
		createdAt: time.Now(),
	}

	printUser(userInstance)

}

func printUser(userVar user) {
	fmt.Printf("Hello %s, %s!\n", userVar.firstName, userVar.lastName)
	fmt.Printf("Thank you for adding %s as your email %s.\n", userVar.email, userVar.firstName)
	fmt.Printf("Your id was created at %s.\n", userVar.createdAt.Format(time.RFC3339))
}

func inputChoice() (string, string, string) {
	var firstName, lastName, email string
	fmt.Print("Please enter your first name: ")
	fmt.Scan(&firstName)
	fmt.Print("Please enter your Last name: ")
	fmt.Scan(&lastName)
	fmt.Print("Please enter your email address: ")
	fmt.Scan(&email)
	return firstName, lastName, email
}
