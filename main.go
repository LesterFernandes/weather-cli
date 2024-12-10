package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
)

type Weather struct {
	Location struct {
		Name    string `json:"name"`
		Country string `json:"country"`
	} `json:"location"`
	Current struct {
		Tempc     float64 `json:"temp_c"`
		Condition struct {
			Text string `json:"text"`
		} `json:"condition"`
	} `json:"current"`
	Forecast struct {
		Forecastday []struct {
			Hour []struct {
				TimeEpoch int64   `json:"time_epoch"`
				Tempc     float64 `json:"temp_c"`
				Condition struct {
					Text string `json:"text"`
				} `json:"condition"`
				ChanceOfRain int64 `json:"chance_of_rain"`
			} `json:"hour"`
		} `json:"forecastday"`
	} `json:"forecast"`
}

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	locationCode := "London"
	if len(os.Args) >= 2 {
		locationCode = os.Args[1]
	}

	apiKey := os.Getenv("API_KEY")
	response, err := http.Get("https://api.weatherapi.com/v1/forecast.json?key=" + apiKey + "&q=" + locationCode + "&days=1&aqi=no&alerts=no")

	if err != nil {
		panic(err)
	}

	if response.StatusCode != 200 {
		panic("Weather api not avlbl")
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	var weatherReport Weather
	errah := json.Unmarshal(body, &weatherReport)
	if errah != nil {
		panic(errah)
	}

	location, current, hours := weatherReport.Location, weatherReport.Current, weatherReport.Forecast.Forecastday[0].Hour

	fmt.Printf("%s, %s, %fC, %s \n ", location.Name, location.Country, current.Tempc, current.Condition.Text)

	for _, hour := range hours {
		date := time.Unix(hour.TimeEpoch, 0)
		// fmt.Println(hour)
		if date.Before(time.Now()) {
			continue
		}

		msg := fmt.Sprintf("%s, %s, %fC ", date.Format("15:04"), hour.Condition.Text, hour.Tempc)
		if hour.ChanceOfRain < 40 {
			fmt.Println(msg)
		} else {
			color.Red(msg)
		}
	}
}
