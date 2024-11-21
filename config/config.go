package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	OpenTripMapKey string
	WikiVoyageURL  string
	OSMEndpoint    string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found. Loading defaults...")
	}

	return &Config{
		OpenTripMapKey: os.Getenv("OPENTRIPMAP_KEY"),
		WikiVoyageURL:  "https://en.wikivoyage.org/w/api.php",
		OSMEndpoint:    "https://overpass-api.de/api/interpreter",
	}
}
