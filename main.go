package main

import (
	"fmt"
	"strings"
)

func main() {

	// variables
	conferenceName := "Go Conference" //syntantic sugar
	const conferenceTickets = 50
	var remainingTickets uint = 50
	//using slice for storing tickets instead of array
	bookings := []string{}

	//calling the  greetUsers function:
	greetUsers(conferenceName, conferenceTickets, remainingTickets)

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

		//calling validateUserInput func:
	    isValidName, isValidEmail, isValidTicketNumber:=validateUserInput(userFirstname,userLastname,email,userTickets,remainingTickets )

		//check if the user is booking more tickets than we have in total
		if isValidName && isValidName && isValidTicketNumber {
			//logic for updating the number of tickets remaining:
			remainingTickets = remainingTickets - userTickets

			//logic for updating the user's first and last name:
			bookings = append(bookings, userFirstname+" "+userLastname)

			fmt.Printf("Thank you %v %v for booking %v tickets, you will receive your confirmation for the tickets at %v\n ", userFirstname, userLastname, userTickets, email)
			fmt.Printf("The remaining tickets are now %v\n", remainingTickets)

			//call the function printFirstNames:
			firstNames := getFirstNames(bookings)
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

func greetUsers(confName string, confTickets int, remainingTickets uint) {
	//greetings/welcome messages
	fmt.Printf("welcome to our %v booking application\n", confName)
	fmt.Printf("we have  %v tickets in total and we now have %v tickets available\n", confTickets, remainingTickets)
	fmt.Println("Get your tickets here to now!")
}

//[] indicates that we are returning a slice of strings
func getFirstNames(bookings []string) []string {
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

func validateUserInput(userFirstname string , userLastname string , email string,userTickets uint, remainingTickets uint ) (bool,bool,bool){
	//validating user input:
		isValidName := len(userFirstname) > 2 && len(userLastname) > 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
		return isValidName, isValidEmail, isValidTicketNumber
	
}
