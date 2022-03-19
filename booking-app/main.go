package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets = 50 //constant
	var remainingTickets uint = 50
	//var bookings [50]string //array with fixed size
	var bookings []string //slice with dynamic size

	fmt.Printf("Welcome to our %v booking application\n", conferenceName)
	fmt.Printf("We have %v tickets in total and there are %v tickets remaining\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	//var bookings = [50]string{} //array

	for {
		var firstName string
		var lastName string
		var userEmail string
		var userTickets uint

		// ask user for their information
		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName) //using a pointer

		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email: ")
		fmt.Scan(&userEmail)

		fmt.Println("How many tickets would you like to order? ")
		fmt.Scan(&userTickets)

		if userTickets > remainingTickets {
			fmt.Printf("We only have %v tickets remaining!\n", remainingTickets)
			continue
		}

		remainingTickets -= userTickets
		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, userEmail)
		fmt.Printf("%v tickets remaing for the %v\n", remainingTickets, conferenceName)

		var firstNames []string
		for _, booking := range bookings { // _ is a variable not being used
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}
		fmt.Printf("These are all our bookings: %v\n", firstNames)

		noTicketsRemaining := remainingTickets == 0
		if noTicketsRemaining {
			fmt.Println("All tickets have been sold! Come back next year!")
			break
		}

	}

}
