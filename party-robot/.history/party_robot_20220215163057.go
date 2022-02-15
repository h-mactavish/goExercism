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
	var (
		line1 string
		line2 string
		line3 string
	)
	line1 = fmt.Sprintf("Welcome to my party, %s!", name)
	line2 = fmt.Sprintf(, table)
	line3 = fmt.Sprintf("You have been assigned to table %0d. Your table is %s,exactly %.2f meters from here.",table, direction, distance)

	//panic("Please implement the AssignTable function")
}
