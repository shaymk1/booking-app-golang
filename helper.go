
package main
import "strings"

//  first letter of the function is capital because we are exporting the function
func ValidateUserInput(userFirstname string, userLastname string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	//validating user input:
	isValidName := len(userFirstname) > 2 && len(userLastname) > 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTicketNumber

}