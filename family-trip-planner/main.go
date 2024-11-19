package main

import (
	"fmt"
	"family-trip-planner/input"
	"family-trip-planner/api"
	"family-trip-planner/itinerary"
)

func main() {
	// Get user input
	userInput := input.GetUserInput()

	// Fetch data from APIs
	places := api.FetchPlaces(userInput.Location, userInput.Days)

	// Generate itinerary
	itinerary := itinerary.GenerateItinerary(userInput, places)

	// Display the itinerary
	fmt.Println("Your Itinerary:")
	for _, day := range itinerary {
		fmt.Printf("Day %d: %s\n", day.Day, day.Plan)
	}
}
