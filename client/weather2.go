package client

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Props struct {
	Forecast       string `json:"forecast"`
	ForecastOffice string `json:"forecastOffice"`
	GridX          int    `json:"gridX"`
	GridY          int    `json:"gridY"`
}

type Weather struct {
	Props Props `json:"properties"`
}

func GetWeather(latitude, longitude float64) (string, error) {
	var err error
	var weather Weather
	// Create a new HTTP client.
	client := &http.Client{}

	// Create a new GET request to a URL that reflects headers.
	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.weather.gov/points/%f,%f", latitude, longitude), nil)
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	// Set the User-Agent header.
	req.Header.Set("User-Agent", "(amandaweatherapp.com, example@email.com)")
	// Perform the request.
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error performing request: %v", err)
	}
	defer resp.Body.Close()

	// Read and print the response body.
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}
	err = json.Unmarshal(body, &weather)
	if err != nil {
		log.Fatalf("Error marshalling response body: %v", err)
	}
	// fmt.Printf("Weather Props: %+v\n", weather)

	office := strings.Split(weather.Props.ForecastOffice, "/")
	forecastDetails, err := GetForecast(ForecastInput{
		Office: office[len(office)-1],
		GridX:  weather.Props.GridX,
		GridY:  weather.Props.GridY,
	})
	if err != nil {
		log.Fatalf("Error getting forecast: %v", err)
	}

	formattedForecast, err := FormatForecastResponse(FormatForecastInput{ForecastData: *forecastDetails})
	if err != nil {
		log.Fatalf("Error formatting forecast response: %v", err)
	}

	result, err := json.Marshal(formattedForecast)
	if err != nil {
		log.Fatalf("Error marshalling final response: %v", err)
	}

	return string(result), nil
}

type ForecastInput struct {
	Office string `json:"office"`
	GridX  int    `json:"gridX"`
	GridY  int    `json:"gridY"`
}
type ForecastPeriods struct {
	Name             string `json:"name"`
	ShortForecast    string `json:"shortForecast"`
	DetailedForecast string `json:"detailedForecast"`
}
type ForecastProperties struct {
	Periods []ForecastPeriods `json:"periods"`
}
type ForecastResponse struct {
	Props ForecastProperties `json:"properties"`
}

func GetForecast(input ForecastInput) (*ForecastResponse, error) {
	var err error

	// Create a new HTTP client.
	client := &http.Client{}

	// Create a new GET request to a URL that reflects headers.
	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.weather.gov/offices/%s/grids/%d,%d/forecast", input.Office, input.GridX, input.GridY), nil)
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	// Set the User-Agent header.
	req.Header.Set("User-Agent", "(amandaweatherapp.com, example@email.com)")
	// Perform the request.
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error performing request: %v", err)
	}
	defer resp.Body.Close()

	// Read and print the response body.
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}
	var forecast ForecastResponse
	err = json.Unmarshal(body, &forecast)
	if err != nil {
		log.Fatalf("Error marshalling response body: %v", err)
	}
	return &forecast, nil
}

func CharacterizeTemperature(tempF int) string {
	if tempF >= 85 {
		return "Hot"
	} else if tempF <= 60 {
		return "Cold"
	} else {
		return "Moderate"
	}
}

type FormatForecastInput struct {
	ForecastData ForecastResponse
}

type FormattedForecastResponse struct {
	Name          string `json:"name"`
	ShortForecast string `json:"shortForecast"`
	TempCharacter string `json:"tempCharacter"`
}

func FormatForecastResponse(input FormatForecastInput) (*FormattedForecastResponse, error) {
	if len(input.ForecastData.Props.Periods) > 0 {
		tempStr := strings.Split(input.ForecastData.Props.Periods[0].DetailedForecast, " ")
		temp, err := strconv.Atoi(tempStr[len(tempStr)-2])
		if err != nil {
			log.Fatalf("Error converting temperature to int: %v", err)
		}
		return &FormattedForecastResponse{
			Name:          input.ForecastData.Props.Periods[0].Name,
			ShortForecast: input.ForecastData.Props.Periods[0].ShortForecast,
			TempCharacter: CharacterizeTemperature(temp), // Placeholder temperature
		}, nil
	}
	return nil, nil
}
