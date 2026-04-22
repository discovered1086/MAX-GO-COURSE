package main

import "fmt"

func main() {
	age := 10

	//var agePointer *int
	//agePointer = &age

	agePointer := &age

	fmt.Println("The address of age is", agePointer)
	fmt.Println("The value of age is", *agePointer)

	//Calling the function by passing the pointer to update the value
	modifyAgeByFiveWithPointer(agePointer)
	fmt.Println("The updated value of age variable is ", age)
}

func modifyAgeByFiveWithPointer(agePointer *int) {
	*agePointer = *agePointer + 5
}
