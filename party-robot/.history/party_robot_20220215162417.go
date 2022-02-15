package partyrobot

import "fmt"

// Welcome greets a person by name.
func Welcome(name string) string {
	return "Welcome to my party, " + name
	//panic("Please implement the Welcome function")
}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {
	result := fmt.Sprintf("Happy birthday %s! You are now %d years old!", name, age)
	return result
	//panic("Please implement the HappyBirthday function")
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
	line1 := fmt.Sprint("Welcome to my party, %s!")
	//panic("Please implement the AssignTable function")
}
