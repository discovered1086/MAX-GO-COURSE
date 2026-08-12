package userPackage

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	FirstName string
	lastName  string
	email     string
	createdAt time.Time
}

type AdminUser struct {
	Designation string
	Department  string
	User
}

func NewAdminUser(firstName, lastName, email, department, designation string) AdminUser {
	return AdminUser{
		Designation: designation,
		Department:  department,
		User: User{
			FirstName: firstName,
			lastName:  lastName,
			email:     email,
		},
	}
}

func (userVar *User) PrintUserMethod() {
	fmt.Printf("Hello %s, %s!\n", userVar.FirstName, userVar.lastName)
	fmt.Printf("Thank you for adding %s as your email %s.\n", userVar.email, userVar.FirstName)
	fmt.Printf("Your id was created at %s.\n", userVar.createdAt.Format(time.RFC3339))
}

func (userVar *User) ChangeUserDetails() {
	userVar.FirstName = ""
	userVar.lastName = ""
	userVar.email = ""
}

// NewUserDetailsWithMethod Constructor function
func NewUserDetailsWithMethod(firstName, lastName, email string) (*User, error) {
	if firstName == "" || lastName == "" || email == "" {
		return nil, errors.New("invalid input")
	}
	return &User{
		FirstName: firstName,
		lastName:  lastName,
		email:     email,
		createdAt: time.Now(),
	}, nil
}
