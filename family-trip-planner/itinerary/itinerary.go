package itinerary

import (
	"fmt"
	"family-trip-planner/api"
	"family-trip-planner/input"
)

// DayPlan represents the plan for a single day
type DayPlan struct {
	Day  int
	Plan string
}

// GenerateItinerary creates a plan based on places and user input
func GenerateItinerary(userInput input.UserInput, places []api.Place) []DayPlan {
	var itinerary []DayPlan
	for i := 0; i < userInput.Days; i++ {
		if i < len(places) {
			itinerary = append(itinerary, DayPlan{
				Day:  i + 1,
				Plan: fmt.Sprintf("Visit %s - %s", places[i].Name, places[i].Description),
			})
		} else {
			itinerary = append(itinerary, DayPlan{
				Day:  i + 1,
				Plan: "Explore the city freely!",
			})
		}
	}
	return itinerary
}
