package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	// Information for the location passed in
	locationName := os.Getenv("LOCATION")
	if locationName == "" {
		locationName = "The Cloudflare San Francisco Office"
	}

	// Float64 with default value
	var lat float64 = 37.780259
	if latStr := os.Getenv("LATITUDE"); latStr != "" {
		if parsedLat, err := strconv.ParseFloat(latStr, 64); err == nil {
			lat = parsedLat
		}
	}

	// Float64 with default value
	var lon float64 = -122.390519
	if lonStr := os.Getenv("LONGITUDE"); lonStr != "" {
		if parsedLon, err := strconv.ParseFloat(lonStr, 64); err == nil {
			lon = parsedLon
		}
	}

	// Build the API URL with current weather parameters
	url := fmt.Sprintf("http://api.open-meteo.com/v1/forecast?latitude=%.6f&longitude=%.6f&current=temperature_2m,apparent_temperature,is_day,precipitation,rain,showers,snowfall,weather_code,cloud_cover&timezone=auto",
		lat, lon)

	// Make the HTTP request
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error making request: %v\n", err)
		return
	}
	defer resp.Body.Close()

	// Check if the request was successful
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("API request failed with status: %d\n", resp.StatusCode)
		body, _ := io.ReadAll(resp.Body)
		fmt.Printf("Response: %s\n", string(body))
		return
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response: %v\n", err)
		return
	}

	// Parse the JSON response
	var weather WeatherResponse
	err = json.Unmarshal(body, &weather)
	if err != nil {
		fmt.Printf("Error parsing JSON: %v\n", err)
		return
	}

	// Parse the time
	currentTime, err := time.Parse("2006-01-02T15:04", weather.Current.Time)
	if err != nil {
		currentTime = time.Now()
	}

	// Display the weather information
	fmt.Printf("Weather at %s at: %s\n", locationName, currentTime.Format("2006-01-02 15:04"))
	fmt.Printf("Located at %.6f at: %.6f\n", lat, lon)
	fmt.Println("----------------------------------------")

	fmt.Printf("Condition: %s\n", GetWeatherDescription(weather.Current.WeatherCode))
	fmt.Printf("Temperature: %.1f%s\n", weather.Current.Temperature2m, weather.CurrentUnits.Temperature2m)
	fmt.Printf("Cloud cover: %d%s\n", weather.Current.CloudCover, weather.CurrentUnits.CloudCover)
	if weather.Current.Rain > 0 {
		fmt.Printf("Rain: %.1f %s\n", weather.Current.Rain, weather.CurrentUnits.Rain)
	}

	if weather.Current.Snowfall > 0 {
		fmt.Printf("Snow: %.1f %s\n", weather.Current.Snowfall, weather.CurrentUnits.Snowfall)
	}
}
