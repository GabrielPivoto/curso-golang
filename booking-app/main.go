package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50 //constant
	var remainingTickets = 50

	fmt.Printf("Welcome to our %v booking application\n", conferenceName)
	fmt.Printf("We have %v tickets in total and there are %v tickets remaining\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}
