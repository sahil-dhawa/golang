package main

import (
	"booking-project/helper"
	"fmt"
)

func main() {

	fmt.Println("Hello world")
	var conferenceName = "Go Conference"
	var userName string
	var totalTickets uint
	var pendingTickets uint
	var firstName string
	var lastName string
	var ticketsBooked uint
	var bookingSlots [50]string
	var totalSum int

	helper.GreetUsers(conferenceName)

	totalSum = helper.CalculateSum(10, 20)

	fmt.Printf("Sum is %v\n", totalSum)

	totalTickets = 50
	fmt.Println("Please enter firstName")
	fmt.Scan(&firstName)

	fmt.Println("Please enter lastName")
	fmt.Scan(&lastName)

	fmt.Println("Please enter ticketsBooked")
	fmt.Scan(&ticketsBooked)

	fmt.Println("Please enter username")
	fmt.Scan(&userName)

	fmt.Println("Welcome to", conferenceName, "world")
	fmt.Printf("The username is %v with tickets booked %v is\n", userName, ticketsBooked)

	pendingTickets = totalTickets - ticketsBooked

	fmt.Printf("Total pending Tickets are %v", pendingTickets)

	fmt.Printf("Total number of slots is  %v", bookingSlots)
}
