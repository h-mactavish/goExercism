package techpalace

import "strings"

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace, " + customer
	//panic("Please implement the WelcomeMessage() function")
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	var deco string = strings.Repeat("*", numStarsPerLine)
	var result string = deco + "\n" + welcomeMsg + "\n" + deco
	return result
	//panic("Please implement the AddBorder() function")
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	var result string = strings.TrimPrefix("*", oldMsg)
	result = strings.TrimPrefix("`", result)
	return result
	//panic("Please implement the CleanupMessage() function")
}
