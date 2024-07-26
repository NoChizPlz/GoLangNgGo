package main

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"io"
	"net/http"
	"os"
	"time"
)

type Weather struct {
	Location struct {
		Name    string `json:"name"`
		Country string `json:"country"`
	} `json:"location"`
	Current struct {
		TempC     float64 `json:"temp_c"`
		Condition struct {
			Text string `json:"text"`
		}
	} `json:"current"`
	Forecast struct {
		ForecastDay []struct {
			Hour []struct {
				TimeEpoch int64   `json:"time_epoch"`
				TempC     float64 `json:"temp_c"`
				Condition struct {
					Text string `json:"text"`
				} `json:"condition"`
				ChanceOfRain float64 `json:"chance_of_rain"`
			} `json:"hour"`
		} `json:"forecastday"`
	}
}

func main() {
	city := "Manila"

	if len(os.Args) >= 2 {
		city = os.Args[1]
	}
	res, err := http.Get("https://api.weatherapi.com/v1/forecast.json?key=c1f745c08f8442ef8b654349242607&q=" + city + "&days=1&aqi=no&alerts=no")

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		panic(res.Body)
	}

	body, err := io.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}

	var weather Weather
	err = json.Unmarshal(body, &weather)
	if err != nil {
		panic(err)
	}

	location, current, hours := weather.Location, weather.Current, weather.Forecast.ForecastDay[0].Hour

	fmt.Printf(
		"Location: %s, Country: %s, Temp: %.2f, Condition: %s\n",
		location.Name,
		location.Country,
		current.TempC,
		current.Condition.Text,
	)

	for _, hours := range hours {
		date := time.Unix(hours.TimeEpoch, 0)
		if date.Before(time.Now()) {
			continue
		}
		forecast := fmt.Sprintf("%s -  %.2fC - %.f%% - %s\n",
			date.Format("Mon 02-Jan-06 03:04 PM"),
			hours.TempC,
			hours.ChanceOfRain,
			hours.Condition.Text,
		)

		if hours.ChanceOfRain < 60 {
			fmt.Print(forecast)
		} else {
			color.Red(forecast)
		}
	}
}
