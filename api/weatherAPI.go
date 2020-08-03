package api

import (
	env_vars "api-test/env-vars"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var myHttpClient = &http.Client{Timeout: 10 * time.Second}

type WeatherObject struct {
	Coord      Coord
	Weather    []Weather
	Main       Main
	Visibility int
	Wind       Wind
	Timestamp string
}

func (r WeatherObject) Error() string {
	panic("implement me")
}

type Wind struct {
	Speed float32
	Deg   float32
}
type Weather struct {
	Id          int
	Main        string
	Description string
	Icon        string
}
type Main struct {
	Temp      float32
	FeelsLike float32 `json:"feels_like"`
	TempMin   float32 `json:"temp_min"`
	TempMax   float32 `json:"temp_max"`
	Pressure  float32
	Humidity  float32
}
type Coord struct {
	Lon float32 `json:"lon"`
	Lat float32 `json:"lat"`
}

func GetWeather(location string) (*WeatherObject, error) {
	apiKey:= env_vars.EnvVariable("weather_api_key")
	resp, err := myHttpClient.Get("http://api.openweathermap.org/data/2.5/weather?q=" + location + "&appid=" + apiKey + "&units=metric")
	var responseJson WeatherObject
	if err != nil {
		print(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	errUnmarshall := json.Unmarshal(body, &responseJson)
	if errUnmarshall != nil {
		fmt.Printf("There was an error decoding the json. err = %s", err)
		return nil, errors.New("error decoding JSON")
	}
	currentTime := time.Now()
	responseJson.Timestamp =currentTime.Format("15:04:05, 02-01-2006")
	return &responseJson, nil
}
