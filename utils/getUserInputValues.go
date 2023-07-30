package utils

import (
	"fmt"
)

func GetUserInputvalues() (string, string, string, string, uint) {
	var userName string
	var firstName string
	var lastName string
	var email string
	var userTicket uint

	// asks for user name
	fmt.Println("Enter your  username: ")
	fmt.Scan(&userName)

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email:")
	fmt.Scan(&email)

	fmt.Println("Enter your number of tickets:")
	fmt.Scan(&userTicket)

	return userName, firstName, lastName, email, userTicket
}