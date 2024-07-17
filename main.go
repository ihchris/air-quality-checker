package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// AirVisualResponse represents the structure of the response from AirVisual API
type AirVisualResponse struct {
	Status string `json:"status"`
	Data   struct {
		City    string `json:"city"`
		State   string `json:"state"`
		Country string `json:"country"`
		Current struct {
			Pollution struct {
				Aqius  int    `json:"aqius"`
				Mainus string `json:"mainus"`
			} `json:"pollution"`
		} `json:"current"`
	} `json:"data"`
}

// fetchAirQuality fetches the air quality data from AirVisual API
func fetchAirQuality(apiKey, city, state, country string) (*AirVisualResponse, error) {
	url := fmt.Sprintf("https://api.airvisual.com/v2/city?city=%s&state=%s&country=%s&key=%s", city, state, country, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var airQuality AirVisualResponse
	if err := json.NewDecoder(resp.Body).Decode(&airQuality); err != nil {
		return nil, err
	}

	return &airQuality, nil
}

func main() {
	apiKey := "your_airvisual_api_key" // Replace with your AirVisual API key
	city := "Los Angeles"              // Replace with the city you want to check
	state := "California"              // Replace with the state you want to check
	country := "USA"                   // Replace with the country you want to check

	airQuality, err := fetchAirQuality(apiKey, city, state, country)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error fetching air quality data: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Air Quality in %s, %s, %s:\n", airQuality.Data.City, airQuality.Data.State, airQuality.Data.Country)
	fmt.Printf("AQI (US): %d\n", airQuality.Data.Current.Pollution.Aqius)
	fmt.Printf("Main Pollutant: %s\n", airQuality.Data.Current.Pollution.Mainus)
}
