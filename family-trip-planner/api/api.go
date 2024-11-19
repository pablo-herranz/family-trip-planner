package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

API_KEY = "5ae2e3f221c38a28845f05b63e95772af969fc4aa7e42b827ccb9e29"

// Place represents a place of interest
type Place struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// FetchPlaces queries an API for places of interest
func FetchPlaces(location string, days int) []Place {
	// Replace with a real API URL
	apiURL := fmt.Sprintf("https://api.example.com/places?location=%s&days=%d", location, days)

	// Make the API request
	resp, err := http.Get(apiURL)
	if err != nil {
		fmt.Println("Error fetching places:", err)
		return nil
	}
	defer resp.Body.Close()

	// Parse the JSON response
	body, _ := ioutil.ReadAll(resp.Body)
	var places []Place
	json.Unmarshal(body, &places)

	return places
}
