package config

import (
    "log"
    "os"

    "github.com/joho/godotenv"
)

func LoadAPIKey() string {
    err := godotenv.Load()
    if err != nil {
        log.Println("Warning: .env file not found.")
    }
    apiKey := os.Getenv("OPENTRIPMAP_KEY")
    if apiKey == "" {
        log.Fatal("API key not set in .env")
    }
    return apiKey
}
