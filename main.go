package main

import "fmt"

func main() {

	// variables
	conferenceName := "Go Conference" //syntantic sugar
	const conferenceTickets = 50
	var remainingTickets uint = 50
	//using slice for storing tickets instead of array
	bookings := []string{}

	//grretings/welcome messages
	fmt.Printf("welcome to our %v booking app\n", conferenceName) //f is like f string in pyhton, for formating
	fmt.Printf("we have  %v tickets in total and we now have %v tickets available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to now!")

	// infinite loop simulating different users booking
	for {
		// ticket-booking
		//ask the user to input the username
		var userFirstname string
		var userLastname string
		var email string
		var userTickets uint

		fmt.Println("Enter your first name:  ")
		//a function to point at the data stored in memory->pointer
		fmt.Scan(&userFirstname)

		fmt.Println("Enter your first last name:  ")
		//a function to point at the data stored in memory->pointer
		fmt.Scan(&userLastname)

		fmt.Println("Enter your email:  ")
		//a function to point at the data stored in memory->pointer
		fmt.Scan(&email)

		fmt.Println("How many tickets you want:  ")
		//a function to point at the data stored in memory->pointer
		fmt.Scan(&userTickets)

		//logic for updating the number of tickets:
		remainingTickets = remainingTickets - userTickets

		//logic for updating the user's first and last name:
		bookings = append(bookings, userFirstname+" "+userLastname)

		fmt.Printf("Thank you %v %v for booking %v tickets, you will receive your confirmation for the tickets at %v\n ", userFirstname, userLastname, userTickets, email)
		fmt.Printf("The remaining tickets are now %v\n", remainingTickets)
		fmt.Printf("This is how many %v we have\n", bookings)

	}

}
