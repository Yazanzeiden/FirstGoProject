package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	bookings := []string{}
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v conference Tickets tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	for {

		var firstName string
		var lastName string
		var email string
		var userTicktes int
		//ask your for thier name
		fmt.Println("Enter your first name:")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name:")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email:")
		fmt.Scan(&email)

		fmt.Println("Enter number of tickets:")
		fmt.Scan(&userTicktes)

		if userTicktes <= int(remainingTickets) {

			remainingTickets = remainingTickets - uint(userTicktes)
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank you %v %v for booking %v tickets. you will receive a confirmation email at %v\n", firstName, lastName, userTicktes, email)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}

			fmt.Printf("the first names of booking are : %v\n", firstNames)

			if remainingTickets == 0 {
				// end program
				fmt.Println("Our conference is booked out. comeback next year")
				break
			}
		} else {
			fmt.Printf("we only have  %v tickets remaining, so you cant book %v tickets\n", remainingTickets, userTicktes)

		}

	}
}
