package main

import (
    "log"

    "github.com/pablo-herranz/family-trip-planner/config"
    "github.com/pablo-herranz/family-trip-planner/internal/api"
    "github.com/pablo-herranz/family-trip-planner/internal/planner"
)

func main() {
    // Load API key
    apiKey := config.LoadAPIKey()

    // Initialize OpenTripMap client
    client := api.NewOpenTripMapClient(apiKey)

    // Example: Kuala Lumpur coordinates
    pois, err := client.GetPOIs(3.1390, 101.6869, 1000)
    if err != nil {
        log.Fatalf("Error fetching POIs: %v", err)
    }

    // Generate a simple itinerary
    planner.GenerateItinerary(pois, 5)
}
