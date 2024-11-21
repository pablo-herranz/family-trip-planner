package itinerary

import "github.com/yourusername/family-trip-planner/internal/api"

type Planner struct {
	TripMapClient *api.OpenTripMapClient
}

func NewPlanner(client *api.OpenTripMapClient) *Planner {
	return &Planner{TripMapClient: client}
}

func (p *Planner) PlanTrip(destination string, duration int, kidsAges []int) (map[string]interface{}, error) {
	// Example: Fetch POIs around the destination
	pois, err := p.TripMapClient.GetPOIsByRadius(3.1390, 101.6869, 1000) // Example coordinates for Kuala Lumpur
	if err != nil {
		return nil, err
	}

	// Placeholder for trip plan logic
	tripPlan := map[string]interface{}{
		"destination": destination,
		"pois":        pois,
		"duration":    duration,
		"kidsAges":    kidsAges,
	}

	return tripPlan, nil
}
