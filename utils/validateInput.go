package utils

import (
	"strings"
)

func ValidateInputs(firstName string, lastName string, email string, userTicket uint, remainingTickets uint) (bool, bool, bool) {
	isNameValid := len(firstName) >= 2 && len(lastName) >= 2
	isEmailValid := strings.Contains(email, "@")
	isTicketNumberValid := userTicket > 0 && userTicket <= remainingTickets
	return isEmailValid, isNameValid, isTicketNumberValid
}
