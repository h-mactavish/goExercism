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
	var line string
	fmt.Sprintf("Welcome to my party, %s!", line)
	fmt.Sprintf("You have been assigned to table %0d.", line)
	//panic("Please implement the AssignTable function")
}