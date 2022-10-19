package main

import (
	"fmt"
	"strings"
)

// package level variables
const conferenceTickets = 50

var conferenceName string = "Go Conference"
var remainingTickets uint = 50
var bookings = []string{} //using slice for storing tickets instead of array

func main() {

	//calling the  greetUsers function:
	greetUsers()

	// infinite loop simulating different users booking
	for {
		// callling getUserInput function:
		userFirstname, userLastname, email, userTickets := getUserInput()

		//calling validateUserInput func:
		isValidName, isValidEmail, isValidTicketNumber := ValidateUserInput(userFirstname, userLastname, email, userTickets, remainingTickets)

		//check if the user is booking more tickets than we have in total
		if isValidName && isValidEmail && isValidTicketNumber {

			//calling bookingTickets function:
			bookingTickets(userTickets, userFirstname, userLastname, email)

			//call the function printFirstNames:
			firstNames := getFirstNames()
			fmt.Printf("The first names are now %v\n", firstNames)

			//check if all tickets are booked and no tickets are left:
			if remainingTickets == 0 {
				//end program
				fmt.Println("Sorry , all our tickets are booked out")
				break
			}

			// } else {
			// 	//end program
			// 	fmt.Printf("Sorry, we only have %v tickets available, so you can't book %v tickets\n ", remainingTickets, userTickets)
			// 	break}

		} else {
			if !isValidName {
				fmt.Println("Sorry, the first name or last name is too short, it needs to be more than 2 characters")

			}
			if !isValidEmail {
				fmt.Println("Sorry, the email address is incorrect, email address must contain an @ character")

			}
			if !isValidTicketNumber {
				fmt.Println("Sorry, the ticket number is incorrect, please enter a valid ticket number")

			}

		}

	}

}

func greetUsers() {
	//greetings/welcome messages
	fmt.Printf("welcome to our %v booking application\n", conferenceName)
	fmt.Printf("we have  %v tickets in total and we now have %v tickets available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to now!")
}

// [] indicates that we are returning a slice of strings
func getFirstNames() []string {
	//iterating through the first names only from bookings"
	firstNames := []string{}

	//using _ for the index , because the first we dnt want to get error for the unused variables
	for _, booking := range bookings {
		//strings split strings with whitespace separator and slice it into 2 elements
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames

}
func getUserInput() (string, string, string, uint) {
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

	return userFirstname, userLastname, email, userTickets

}

func bookingTickets(userTickets uint, userFirstname string, userLastname string, email string) {
	//logic for updating the number of tickets remaining:
	remainingTickets = remainingTickets - userTickets

	//logic for updating the user's first and last name:
	bookings = append(bookings, userFirstname+" "+userLastname)

	fmt.Printf("Thank you %v %v for booking %v tickets, you will receive your confirmation for the tickets at %v\n ", userFirstname, userLastname, userTickets, email)
	fmt.Printf("The remaining tickets are now %v\n", remainingTickets)

}
