package main

import "fmt"

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v conferenceTickets tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var userName string
	var userTicktes int
	//ask your for thier name
	userName = "Tom"
	userTicktes = 2
	fmt.Printf("User %v booked %v tickets.\n", userName, userTicktes)

}
