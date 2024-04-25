package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

// Check credit card number validity in isValidCreditCardNumber
func isValidCreditCardNumber(cardNumber string) bool {
	// Define the regex pattern to match valid credit card numbers
	pattern := `^(?:4|5|6)\d{3}(-?\d{4}){3}$`

	// Compile the regex pattern
	regex := regexp.MustCompile(pattern)

	// Check if the credit card number matches the pattern
	if !regex.MatchString(cardNumber) {
		return false
	}

	// Remove hyphens from the credit card number
	filteredCardNumber := removeHyphens(cardNumber)

	// Check for consecutive repeated digits
	for i := 0; i < len(filteredCardNumber)-3; i++ {
		if filteredCardNumber[i] == filteredCardNumber[i+1] && filteredCardNumber[i] == filteredCardNumber[i+2] && filteredCardNumber[i] == filteredCardNumber[i+3] {
			return false
		}
	}

	return true
}

// Remove hyphens from the strings in removeHyphens
func removeHyphens(s string) string {
	return regexp.MustCompile("-").ReplaceAllString(s, "")
}

func main() {
	// Read user input from terminal
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter a credit card number (or type 'exit' to quit):")

	for scanner.Scan() {
		input := scanner.Text()
		// Exit if user types 'exit'
		if input == "exit" {
			break
		}

		// Check the validity of the credit card number
		isValid := isValidCreditCardNumber(input)
		validity := "Invalid"
		if isValid {
			validity = "Valid"
		}
		fmt.Printf("Credit Card Number: %s : %s\n", input, validity)
		fmt.Println("Enter another credit card number (or type 'exit' to quit):")
	}

	if scanner.Err() != nil {
		fmt.Println("Error reading input:", scanner.Err())
	}
}
