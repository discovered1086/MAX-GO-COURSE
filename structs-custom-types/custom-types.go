package main

import (
	"fmt"
	"strings"
)

type customString string

func (customSting customString) logString() {
	fmt.Println("The custom string is", customSting)
}

func (customSting customString) makeUpperCase() customString {
	return customString(strings.ToUpper(string(customSting)))
}

func main() {
	var custStr customString = "hello world"
	custStr = custStr.makeUpperCase()
	custStr.logString()
}
