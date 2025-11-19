package client

// type WeatherProperties struct {
// 	Forecast       string `json:"forecast"`
// 	ForecastOffice string `json:"forecastOffice"`
// 	GridX          int    `json:"gridX"`
// 	GridY          int    `json:"gridY"`
// }
// type WeatherData struct {
// 	Props WeatherProperties `json:"properties"`
// }

// type WeatherResponse struct {
// 	Name          string `json:"name"`
// 	ShortForecast string `json:"shortForecast"`
// 	TempCharacter string `json:"tempCharacter"`
// }

// func GetWeatherData(latitude, longitude float64) (*WeatherResponse, error) {
// 	// Implementation to call the National Weather Service API
// 	// and return the day(name), short forecast and temperature characterization
// 	var err error
// 	var weatherData WeatherData
// 	// Create a new HTTP client.
// 	client := &http.Client{}

// 	// Create a new GET request to a URL that reflects headers.
// 	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.weather.gov/points/%f,%f", latitude, longitude), nil)
// 	if err != nil {
// 		log.Fatalf("Error creating request: %v", err)
// 	}

// 	// Set the User-Agent header.
// 	req.Header.Set("User-Agent", "(amandaweatherapp.com, example@email.com)")

// 	// Perform the request.
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		log.Fatalf("Error performing request: %v", err)
// 	}
// 	defer resp.Body.Close()

// 	// Read and print the response body.
// 	body, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		log.Fatalf("Error reading response body: %v", err)
// 	}
// 	err = json.Unmarshal(body, &weatherData)
// 	if err != nil {
// 		log.Fatalf("Error marshalling response body: %v", err)
// 	}

// 	office := strings.Split(weatherData.Props.ForecastOffice, "/")

// 	//Call GetForecast to get the short forecast and detailed forecast for temp characterization
// 	details, err := GetForecast(office[len(office)-1], weatherData.Props.GridX, weatherData.Props.GridY)
// 	if err != nil {
// 		log.Fatalf("Error getting forecast: %v", err)
// 	}

// 	//Call GetForecastDetails to parse the forecast details
// 	forecastFinal, err := GetForecastDetails(details)
// 	if err != nil {
// 		log.Fatalf("Error getting forecast details: %v", err)
// 	}

// 	return &WeatherResponse{
// 		Name:          forecastFinal.Name,
// 		ShortForecast: details.Properties.Periods[0].ShortForecast,
// 		TempCharacter: forecastFinal.TempCharacter,
// 	}, nil
// }

// type ForcastProperties struct {
// 	Periods []ForecastPeriods `json:"periods"`
// }

// type ForecastPeriods struct {
// 	Name             string `json:"name"`
// 	ShortForecast    string `json:"shortForecast"`
// 	DetailedForecast string `json:"detailedForecast"`
// }

// type ForecastResponse struct {
// 	Properties ForcastProperties `json:"properties"`
// }

// func GetForecast(office string, gridx, gridy int) (*ForecastResponse, error) {
// 	// Implementation to call the National Weather Service API
// 	// and return the short forecast
// 	var err error
// 	var forecastData ForecastResponse
// 	// Create a new HTTP client.
// 	client := &http.Client{}

// 	// Create a new GET request to a URL
// 	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.weather.gov/gridpoints/%s/%d,%d/forecast", office, gridx, gridy), nil)
// 	if err != nil {
// 		log.Fatalf("Error creating request: %v", err)
// 	}

// 	// Set the User-Agent header.
// 	// This is according to documentation on feature flags and content
// 	req.Header.Set("User-Agent", "(amandaweatherapp.com, example@email.com)")
// 	req.Header.Set("Accept", "application/geo+json")
// 	req.Header.Set("Feature-Flags", "forecast_temperature_qv")

// 	// Perform the request.
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		log.Fatalf("Error performing request: %v", err)
// 	}
// 	defer resp.Body.Close()

// 	// Read and print the response body.
// 	// containing the User-Agent string sent in the request.
// 	body, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		log.Fatalf("Error reading response body: %v", err)
// 	}

// 	err = json.Unmarshal(body, &forecastData)
// 	if err != nil {
// 		log.Fatalf("Error marshalling response body: %v", err)
// 	}
// 	return &forecastData, nil
// }

// func CharacterizeTemperature(tempF int) string {
// 	if tempF >= 85 {
// 		return "Hot"
// 	} else if tempF <= 60 {
// 		return "Cold"
// 	} else {
// 		return "Moderate"
// 	}
// }

// type GetShortForecastResponse struct {
// 	Name          string `json:"name"`
// 	ShortForecast string `json:"shortForecast"`
// 	TempCharacter string `json:"tempCharacter"`
// }

// func GetForecastDetails(forecastData *ForecastResponse) (*GetShortForecastResponse, error) {

// 	if len(forecastData.Properties.Periods) > 0 {
// 		tempStr := strings.Split(forecastData.Properties.Periods[0].DetailedForecast, " ")
// 		temp, err := strconv.Atoi(tempStr[len(tempStr)-2])
// 		if err != nil {
// 			log.Fatalf("Error converting temperature to int: %v", err)
// 		}
// 		return &GetShortForecastResponse{
// 			Name:          forecastData.Properties.Periods[0].Name,
// 			ShortForecast: forecastData.Properties.Periods[0].ShortForecast,
// 			TempCharacter: CharacterizeTemperature(temp), // Placeholder temperature
// 		}, nil
// 	}
// 	return &GetShortForecastResponse{
// 		Name:          "N/A",
// 		ShortForecast: "N/A",
// 		TempCharacter: "N/A",
// 	}, nil
// }

// If I had more time, I would implement error handling for various edge cases, such as invalid coordinates, API rate limiting, and network issues.
// Additionally, I would write unit tests for each function to ensure reliability and correctness.
// Finally, I would consider implementing caching mechanisms to reduce redundant API calls for frequently requested locations.
