package main

import (
	"fmt"
	"log"

	"github.com/yourusername/family-trip-planner/config"
	"github.com/yourusername/family-trip-planner/internal/api"
	"github.com/yourusername/family-trip-planner/internal/itinerary"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize API client
	tripMapClient := api.NewOpenTripMapClient(cfg.OpenTripMapKey)

	// Initialize planner
	planner := itinerary.NewPlanner(tripMapClient)

	// Plan a trip
	plan, err := planner.PlanTrip("Malaysia", 23, []int{3, 5})
	if err != nil {
		log.Fatalf("Failed to plan trip: %v", err)
	}

	fmt.Printf("Your trip plan: %+v\n", plan)
}
