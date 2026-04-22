package main

import "fmt"

func main() {
	age := 10

	fmt.Println("Before function call", age)
	modifyAgeByFive(age)
	fmt.Println("After function call", age)
}

func modifyAgeByFive(age int) {
	age += 5
}
