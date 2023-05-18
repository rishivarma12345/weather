package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type WeatherResponse struct {
	Latitude       float64 `json:"latitude"`
	Longitude      float64 `json:"longitude"`
	GenerationTime float64 `json:"generationtime_ms"`
	UTCOffset      int     `json:"utc_offset_seconds"`
	Timezone       string  `json:"timezone"`
	TimezoneAbbr   string  `json:"timezone_abbreviation"`
	Elevation      float64 `json:"elevation"`
	CurrentWeather Weather `json:"current_weather"`
}

type Weather struct {
	Temperature   float64 `json:"temperature"`
	Windspeed     float64 `json:"windspeed"`
	WindDirection float64 `json:"winddirection"`
	WeatherCode   int     `json:"weathercode"`
	IsDay         int     `json:"is_day"`
	Time          string  `json:"time"`
}

func getWeather(latitude, longitude float64) (*WeatherResponse, error) {
	url := fmt.Sprintf("https://api.open-meteo.com/v1/forecast?latitude=%f&longitude=%f&current_weather=true", latitude, longitude)
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var weatherResponse WeatherResponse
	err = json.Unmarshal(body, &weatherResponse)
	if err != nil {
		return nil, err
	}

	return &weatherResponse, nil
}

func main() {
	// Latitude and Longitude values can be given here
	latitude := 52.52
	longitude := 13.419998
	weather, err := getWeather(latitude, longitude)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Weather:", weather)
	fmt.Println("Current Weather:", weather.CurrentWeather)
	fmt.Printf("Temperature: %.2f °C\n", weather.CurrentWeather.Temperature)
	fmt.Printf("Wind Speed: %.2f m/s\n", weather.CurrentWeather.Windspeed)
	fmt.Printf("Wind Direction: %.2f°\n", weather.CurrentWeather.WindDirection)
	fmt.Printf("Weather Code: %d\n", weather.CurrentWeather.WeatherCode)
	fmt.Printf("Is Day: %d\n", weather.CurrentWeather.IsDay)
	fmt.Printf("Time: %s\n", weather.CurrentWeather.Time)
}
