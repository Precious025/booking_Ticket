package main

import (
	"fmt"

	"go-workspace/helper"
)

var conferenceName string = "Go conference"

const conferenceTickets int = 50

var remainingTickets uint = 50
var bookings = make([]userData, 0)


type userData struct {
	firstName string
	lastName string
	emailAddress string
	numberOfTickets uint
}

func main() {

	greetUsers()

	for {

		firstName, lastName, emailAddress, userTickets := getUserInput()
		isValidName, isValidemail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, emailAddress, userTickets, remainingTickets)

		if isValidName && isValidemail && isValidTicketNumber {

			bookInput(firstName, lastName, emailAddress, userTickets)

			
			fmt.Printf("slice Length: %v\n", len(bookings))

			firstNames := printFirst()
			fmt.Printf("These are the first names of the conference attendees: %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("Our Conference has been booked. Come back next year")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("The first and last name you entered is too short")
			}
			if !isValidemail {
				fmt.Println("The email address does not contain @ sign")
			}
			if !isValidTicketNumber {
				fmt.Println("The number you entered is invalid")
			}
		}

	}

}

func greetUsers() {
	fmt.Println("Welcome to", conferenceName, "!")
	fmt.Printf("We have %v tickets and %v tickets remaining\n", conferenceTickets, remainingTickets)
}

func printFirst() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var emailAddress string
	var userTickets uint

	fmt.Println("Enter your First name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your Last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your Email Address: ")
	fmt.Scan(&emailAddress)

	fmt.Println("Enter the ticket ordered: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, emailAddress, userTickets
}

func bookInput(firstName string, lastName string, emailAddress string, userTickets uint) {

	//create a map to collect userdata
	var userData = userData {
		firstName: firstName,
		lastName: lastName,
		emailAddress: emailAddress,
		numberOfTickets: userTickets,
	}
	

	remainingTickets = remainingTickets - userTickets

	bookings = append(bookings, userData)

	fmt.Printf("This is the map of the attendees: %v\n", bookings)
	
	fmt.Printf("Thank you %v %v for purchasing %v tickets, your ticket has been sent to %v.\n", firstName, lastName, userTickets, emailAddress)
	fmt.Printf("%v is the number of ticket remaining.\n", remainingTickets)
	fmt.Printf("slice type: %T\n", bookings)
}
