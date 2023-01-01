package partyrobot

import "fmt"

// Welcome greets a person by name.
func Welcome(name string) string {
	return fmt.Sprintf("Welcome to my party, %s!", name)
}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {
	return fmt.Sprintf("Happy birthday %s! You are now %d years old!", name, age)
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
	message := Welcome(name) + "\n"
	tableLabel := fmt.Sprintf("%03d", table)
	roundedDistance := fmt.Sprintf("%.1f", distance)

	message += fmt.Sprintf("You have been assigned to table %s. Your table is %s, exactly %s meters from here.\n", tableLabel, direction, roundedDistance)
	message += fmt.Sprintf("You will be sitting next to %s.", neighbor)

	return message
}
