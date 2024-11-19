package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Structure to hold the Overpass response
type OverpassResponse struct {
	Elements []struct {
		Type   string  `json:"type"`
		ID     int     `json:"id"`
		Lat    float64 `json:"lat"`
		Lon    float64 `json:"lon"`
		Tags   struct {
			Amenity string `json:"amenity"`
			Name    string `json:"name"`
		} `json:"tags"`
	} `json:"elements"`
}

func main() {
	// URL-encoded Overpass query to fetch restaurants in the bounding box
	url := "https://overpass-api.de/api/interpreter?data=%5Bout%3Ajson%5D%3Bnode%5Bamenity%3Drestaurant%5D%2848.8566%2C2.3522%2C48.8666%2C2.3622%29%3Bout%20body%3B"

	// Send HTTP GET request
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Error fetching data: %v", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}

	// Unmarshal the JSON response into the OverpassResponse structure
	var overpassResponse OverpassResponse
	err = json.Unmarshal(body, &overpassResponse)
	if err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}

	// Print out the results
	for _, element := range overpassResponse.Elements {
		fmt.Printf("ID: %d, Name: %s, Type: %s, Location: (%f, %f)\n", 
			element.ID, element.Tags.Name, element.Type, element.Lat, element.Lon)
	}
}
