package techpalace

import "string"
// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace, " + customer
	//panic("Please implement the WelcomeMessage() function")
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	var deco= 
	var result string =numStarsPerLine*"*"+"\n"+welcomeMsg+"\n"+numStarsPerLine*"*"
	//panic("Please implement the AddBorder() function")
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	
	//panic("Please implement the CleanupMessage() function")
}
