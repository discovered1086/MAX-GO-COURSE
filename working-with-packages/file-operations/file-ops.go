package file_operations

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func WriteFloatToFile(balance float64, filename string) error {
	floatToText := fmt.Sprintf("%.2f", balance) //Convert balance to a string

	//Permission 0644 means owner of the file can read from or write to the file
	err := os.WriteFile(filename, []byte(floatToText), 0644)

	if err != nil {
		return err
	}
	return nil
}

func ReadFloatFromFile(filename string) (float64, error) {
	bytes, err := os.ReadFile(filename)

	if err != nil {
		return 0.0, errors.New("failed to read float value from a file")
	}

	valueString := string(bytes)
	value, err := strconv.ParseFloat(valueString, 64)

	if err != nil {
		return 0.0, errors.New("failed to parse float from a file")
	}
	return value, nil
}
