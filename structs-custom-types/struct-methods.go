package main

import (
	"fmt"
	"time"
)

type userDetailsWithMethod struct {
	firstName string
	lastName  string
	email     string
	createdAt time.Time
}

func (userVar *userDetailsWithMethod) printUserMethod() {
	fmt.Printf("Hello %s, %s!\n", userVar.firstName, userVar.lastName)
	fmt.Printf("Thank you for adding %s as your email %s.\n", userVar.email, userVar.firstName)
	fmt.Printf("Your id was created at %s.\n", userVar.createdAt.Format(time.RFC3339))
}

func (userVar *userDetailsWithMethod) changeUserDetails() {
	userVar.firstName = ""
	userVar.lastName = ""
	userVar.email = ""
}

// Constructor function
func newUserDetailsWithMethod(firstName, lastName, email string) (*userDetailsWithMethod, error) {
	if firstName == "" || lastName == "" || email == "" {
		return nil, fmt.Errorf("invalid input")
	}
	return &userDetailsWithMethod{
		firstName: firstName,
		lastName:  lastName,
		email:     email,
		createdAt: time.Now(),
	}, nil
}

func main() {
	var firstName, lastName, email string
	fmt.Print("Please enter your first name: ")
	fmt.Scan(&firstName)
	fmt.Print("Please enter your Last name: ")
	fmt.Scan(&lastName)
	fmt.Print("Please enter your email address: ")
	fmt.Scan(&email)
	//
	//fmt.Printf("Hello %s, %s!\n", inputFirstName, inputLastName)
	//fmt.Printf("Thank you for adding %s as your email %s.\n", inputEmail, inputFirstName)

	var userInstance *userDetailsWithMethod

	userInstance, err := newUserDetailsWithMethod(firstName, lastName, email)

	if err != nil {
		fmt.Println(err)
		return
	}

	userInstance.printUserMethod()
	userInstance.changeUserDetails()
	userInstance.printUserMethod()
}
