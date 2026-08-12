package main

import (
	"fmt"
	"structs-custom-types/userPackage"
)

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

	var userInstance *userPackage.User

	userInstance, err := userPackage.NewUserDetailsWithMethod(firstName, lastName, email)

	if err != nil {
		fmt.Println(err)
		return
	}

	userInstance.PrintUserMethod()
	userInstance.ChangeUserDetails()
	userInstance.PrintUserMethod()

	fmt.Println("The first name is", userInstance.FirstName)

	admin := userPackage.NewAdminUser(firstName, lastName, email, "Chemistry", "HOD")

	fmt.Printf("The admin user %s belongs to %s department\n",
		admin.FirstName, admin.Department)

	admin.PrintUserMethod()
}
