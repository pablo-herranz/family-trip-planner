package input

import (
	"fmt"
)

// UserInput holds user preferences
type UserInput struct {
	Location string
	Days     int
}

// GetUserInput prompts the user for input
func GetUserInput() UserInput {
	var input UserInput
	fmt.Println("Welcome to the Trip Planner!")
	fmt.Print("Enter your destination: ")
	fmt.Scanln(&input.Location)
	fmt.Print("Enter the number of days: ")
	fmt.Scanln(&input.Days)
	return input
}
