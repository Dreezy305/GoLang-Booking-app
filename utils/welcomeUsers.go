package utils

import (
	"fmt"
)

func WelcomeUsers(conferenceName string, conferenceTicket int, remainingTickets uint) (string, int, uint) {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTicket, remainingTickets)
	fmt.Println("Get your tickets here")
	return conferenceName, conferenceTicket, remainingTickets
}