package helper

import (
	"fmt"
	"strings"
)

func GetUserInput(userFirstName *string, userLastName *string, userEmail *string) bool {

	fmt.Println("Enter your first name: ")
	fmt.Scan(userFirstName)

	if len(*userFirstName) <= 1 {
		fmt.Println("Please enter a valid first name.")
		return false
	}

	fmt.Println("Enter your last name: ")
	fmt.Scan(userLastName)

	if len(*userLastName) <= 1 {
		fmt.Println("Please enter a valid last name.")
		return false
	}

	fmt.Println("Enter your email: ")
	fmt.Scan(userEmail)

	validEmail := strings.Contains(*userEmail, "@")
	if !validEmail {
		fmt.Println("Please enter a valid email.")
		return false
	}

	return true
}
