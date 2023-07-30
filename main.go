package main

import (
	"booking-app/utils"
	"fmt"
	"sync"
)

const conferenceTicket int = 50

var conferenceName string = "Go Conference"
var remainingTickets uint = 50
var bookings = make([]UserInfo, 0)

type UserInfo struct {
	firstName             string
	lastName              string
	email                 string
	numberOfTicketsBought int
}

// WAIT GROUP BY USER TICKET GEN
var wg sync.WaitGroup

func main() {
	utils.WelcomeUsers(conferenceName, conferenceTicket, remainingTickets)

	// asks for user name
	userName, firstName, lastName, email, userTicket := utils.GetUserInputvalues()

	// validate user input
	isEmailValid, isNameValid, isTicketNumberValid := utils.ValidateInputs(firstName, lastName, email, userTicket, remainingTickets)

	if isEmailValid && isTicketNumberValid && isNameValid {

		bookTicket(firstName, lastName, email, userTicket, userName)

		wg.Add(1)
		go utils.SendTicketToEmail(userTicket, firstName, lastName, email)
		// get users first names
		firstNameResult := getFirstName()

		fmt.Printf("the first names in bookings are: %v\n", firstNameResult)

		fmt.Printf("Bookings: %v\n", bookings)

		if remainingTickets == 0 {
			// stop the bookings
			fmt.Println("sold out")
			// break
		}
	} else {
		if !isNameValid {
			fmt.Println("Name is too short")
		}

		if !isEmailValid {
			fmt.Println("Email is not valid")
		}

		if !isTicketNumberValid {
			fmt.Println("tickets number not allowed")
		}
		// continue
	}
	wg.Wait()
}


func getFirstName() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func bookTicket(firstName string, lastName string, email string, userTicket uint, userName string) {
	remainingTickets = remainingTickets - userTicket

	//  map user data to slice
	var userData = UserInfo{
		firstName:             firstName,
		lastName:              lastName,
		email:                 email,
		numberOfTicketsBought: int(userTicket),
	}
	bookings = append(bookings, userData)
	fmt.Printf("bookings: %v\n", bookings)

	fmt.Printf("User %v booked %v ticket(s).\n", userName, userTicket)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email\n", firstName, lastName, userTicket)

	fmt.Printf("We have %v tickets available for this %v.\n", remainingTickets, conferenceName)
}
