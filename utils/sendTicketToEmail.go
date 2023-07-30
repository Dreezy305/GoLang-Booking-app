package utils

import (
	"fmt"
	"bytes"
	"time"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"sync"
)
var wg sync.WaitGroup

func SendTicketToEmail(userTicket uint, firstName string, lastName string, email string) string {
	time.Sleep(10 * time.Second)
	var ticket string = fmt.Sprintf("%v tickets available for this %v %v.\n", userTicket, firstName, lastName)
	// BREVO EMAIL CLIENT
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	apiKey := os.Getenv("API_KEY")
	url := os.Getenv("EMAIL_CLIENT")
	senderEmail := os.Getenv("SENDER_EMAIL")
	payload := []byte(fmt.Sprintf(`{
        "sender": {"name": "go conference", "email": "%s"},
        "to": [{"email": "%s"}],
        "subject": "Ticket Confirmation",
        "htmlContent": "Hello %s %s. We are happy to inform you that your ticket as been successfully booked."
    }`, senderEmail, email, firstName, lastName))

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("api-key", apiKey)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Println("********************************")
	fmt.Printf("Kindly check your email %v for tickets.\n", email)
	fmt.Println("********************************")
	wg.Done()
	return ticket
}