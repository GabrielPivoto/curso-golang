package main

import "fmt"

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets = 50 //constant
	remainingTickets := 50

	fmt.Printf("Welcome to our %v booking application\n", conferenceName)
	fmt.Printf("We have %v tickets in total and there are %v tickets remaining\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var firstName string
	var lastName string
	var email string
	var userTickets int

	// ask user for their name
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName) //using a pointer

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email: ")
	fmt.Scan(&email)

	fmt.Println("How many tickets would you like to order? ")
	fmt.Scan(&userTickets)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v", firstName, lastName, userTickets, email)

}
