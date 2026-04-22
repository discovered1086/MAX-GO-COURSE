package main

import "fmt"

func main() {
	age := 10

	//var agePointer *int
	//agePointer = &age

	agePointer := &age

	fmt.Println("The address of age is", agePointer)
	fmt.Println("The value of age is", *agePointer)

	//Calling the function by passing the pointer to
	//update the value
	fmt.Println("The updated value of age variable is ",
		modifyAgeByWithPointer(agePointer))
}

func modifyAgeByWithPointer(agePointer *int) int {
	return *agePointer + 5
}
