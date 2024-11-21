package api

import (
    "encoding/json"
    "fmt"
    "net/http"
)

type OpenTripMapClient struct {
    APIKey string
}

func NewOpenTripMapClient(apiKey string) *OpenTripMapClient {
    return &OpenTripMapClient{APIKey: apiKey}
}

func (c *OpenTripMapClient) GetPOIs(lat, lon float64, radius int) ([]string, error) {
    url := fmt.Sprintf("https://api.opentripmap.com/0.1/en/places/radius?radius=%d&lon=%.6f&lat=%.6f&apikey=%s", radius, lon, lat, c.APIKey)
    resp, err := http.Get(url)
    if err != nil {
        return nil, fmt.Errorf("failed to fetch POIs: %w", err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("API returned status %d", resp.StatusCode)
    }

    var result struct {
        Features []struct {
            Properties struct {
                Name string `json:"name"`
            } `json:"properties"`
        } `json:"features"`
    }
    err = json.NewDecoder(resp.Body).Decode(&result)
    if err != nil {
        return nil, fmt.Errorf("failed to decode response: %w", err)
    }

    var pois []string
    for _, feature := range result.Features {
        pois = append(pois, feature.Properties.Name)
    }
    return pois, nil
}
