package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type OpenTripMapClient struct {
	APIKey string
}

func NewOpenTripMapClient(apiKey string) *OpenTripMapClient {
	return &OpenTripMapClient{APIKey: apiKey}
}

func (c *OpenTripMapClient) GetPOIsByRadius(lat, lon float64, radius int) ([]interface{}, error) {
	url := fmt.Sprintf("https://api.opentripmap.com/0.1/en/places/radius?radius=%d&lon=%.6f&lat=%.6f&apikey=%s", radius, lon, lat, c.APIKey)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch POIs: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned status: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var pois []interface{}
	err = json.Unmarshal(body, &pois)
	if err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %w", err)
	}

	return pois, nil
}
