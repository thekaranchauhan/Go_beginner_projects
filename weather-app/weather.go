package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/go-resty/resty/v2"
)

const (
	apiKey      = "YOUR_OPENWEATHERMAP_API_KEY"
	apiEndpoint = "https://api.openweathermap.org/data/2.5/weather"
)

type WeatherData struct {
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
	Name string `json:"name"`
}

func main() {
	if apiKey == "YOUR_OPENWEATHERMAP_API_KEY" {
		fmt.Println("Please replace 'YOUR_OPENWEATHERMAP_API_KEY' with your actual API key.")
		os.Exit(1)
	}

	city := "London" // Replace with the desired city

	weatherData, err := getWeatherData(apiKey, city)
	if err != nil {
		log.Fatal("Error fetching weather data:", err)
	}

	displayWeather(weatherData)
}

func getWeatherData(apiKey, city string) (*WeatherData, error) {
	client := resty.New()

	resp, err := client.R().
		SetQueryParams(map[string]string{
			"q":     city,
			"appid": apiKey,
		}).
		Get(apiEndpoint)

	if err != nil {
		return nil, fmt.Errorf("failed to make request: %v", err)
	}

	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode())
	}

	var weatherData WeatherData
	if err := json.Unmarshal(resp.Body(), &weatherData); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %v", err)
	}

	return &weatherData, nil
}

func displayWeather(data *WeatherData) {
	fmt.Printf("Weather in %s:\n", data.Name)
	fmt.Printf("Temperature: %.2f°C\n", data.Main.Temp)
}
