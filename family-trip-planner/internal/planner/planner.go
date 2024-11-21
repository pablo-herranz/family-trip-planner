package planner

import (
    "fmt"
)

func GenerateItinerary(pois []string, duration int) {
    fmt.Println("Your Trip Itinerary:")
    for i, poi := range pois {
        fmt.Printf("Day %d: Visit %s\n", i+1, poi)
        if i+1 == duration {
            break
        }
    }
}
