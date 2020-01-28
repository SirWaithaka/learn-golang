package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

type weatherData struct {
	Name string `json:"name"`
	Main struct {
		Kelvin float64 `json:"temp"`
	} `json:"main"`
}

type darkSkyData struct {
	Timezone  string `json:"timezone"`
	Currently struct {
		Temperature float64 `json:"temperature"`
	} `json:"currently"`
}

func main() {

	http.HandleFunc("/hi", hi)
	http.HandleFunc("/dark-sky/", darkSkyHandler)
	http.HandleFunc("/weather/", openWeatherHandler)
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8080", nil)
}

func darkSkyHandler(w http.ResponseWriter, r *http.Request) {
	api := "https://api.darksky.net/forecast/88b909d1bda65cd1b2c63c6991478d4d/37.8267,-122.4233"

	data, err := darkSkyquery(api)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(data)
}

func openWeatherHandler(w http.ResponseWriter, r *http.Request) {
	city := strings.SplitN(r.URL.Path, "/", 3)[2]

	api := "http://api.openweathermap.org/data/2.5/weather?APPID=737e78226fb8c6724ef6b8706fadaddd&q="

	data, err := query(city, api)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(data)
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello!"))
}

func hi(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("HI Andela!"))
}

func darkSkyquery(api string) (darkSkyData, error) {
	resp, err := http.Get(api)

	if err != nil {
		return darkSkyData{}, err
	}

	defer resp.Body.Close()

	var d darkSkyData

	if err := json.NewDecoder(resp.Body).Decode(&d); err != nil {
		return darkSkyData{}, err
	}

	return d, nil
}

func query(city string, api string) (weatherData, error) {
	resp, err := http.Get(api + city)

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
