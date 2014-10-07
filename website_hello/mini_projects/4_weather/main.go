// main.go file for learning the weather project.
package main

import (
	"encoding/json"
	"fmt"
	"github.com/codegangsta/martini"
	"net/http"
)

type weatherData struct {
	Name string `json:"name"`
	Main struct {
		Kelvin float64 `json:"temp"`
	} `json:"main"`
}

func main() {
	fmt.Println("Awesome Eric!")

	m := martini.Classic()
	m.Get("/weather/:city", func(w http.ResponseWriter, params martini.Params) {
		city := params["city"]

		weatherData, err := query(city)
		if err != nil {
			panic(err)
			return
		}

		// json newencoder into the responseWriter
		json.NewEncoder(w).Encode(weatherData)
	})
	m.Run()
}

func query(city string) (weatherData, error) {
	resp, err := http.Get("http://api.openweathermap.org/data/2.5/weather?q=" + city)
	if err != nil {
		return weatherData{}, err
	}
	defer resp.Body.Close()

	var d weatherData
	if err := json.NewDecoder(resp.Body).Decode(&d); err != nil {
		return weatherData{}, err
	}
	return d, nil
}
