package main

import (
	"fmt"
)

// Function to validate age
func checkAge(age int) error {
	if age < 18 {
		return fmt.Errorf("age %d is too young for this action", age)
	}
	return nil
}

func main() {
	err := checkAge(20)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Age is valid")
	}
}
