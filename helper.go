
package main
import "strings"


func validateUserInput(userFirstname string, userLastname string, email string, userTickets uint) (bool, bool, bool) {
	//validating user input:
	isValidName := len(userFirstname) > 2 && len(userLastname) > 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTicketNumber

}