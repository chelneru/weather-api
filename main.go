package main

import (
	"api-test/dal"
	"api-test/services"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)


type Response struct {
	Value interface{}
	Message string
}
func getWeather(w http.ResponseWriter, r *http.Request) {

	weatherEntry:= dal.RetrieveLatestWeatherEntry()
	if weatherEntry == nil{
		fmt.Println("Unable to retrieve latest weather entry")
		resp := Response{
			Value:   nil,
			Message: "Unable to retrieve latest weather entry",
		}
		json.NewEncoder(w).Encode(resp)

	}
	resp := Response{
		Value:   weatherEntry,
		Message: "success",
	}
	json.NewEncoder(w).Encode(resp)
}

func getHomeWeather(w http.ResponseWriter, r *http.Request) {

	weatherEntry:= dal.RetrieveLatestHomeWeatherEntry()
	if weatherEntry == nil{
		fmt.Println("Unable to retrieve latest home weather entry")
		resp := Response{
			Value:   nil,
			Message: "Unable to retrieve latest home weather entry",
		}
		json.NewEncoder(w).Encode(resp)

	}
	resp := Response{
		Value:   weatherEntry,
		Message: "success",
	}
	json.NewEncoder(w).Encode(resp)
}
func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.Use(commonMiddleware)
	myRouter.HandleFunc("/weather/{location}", getWeather)
	myRouter.HandleFunc("/home_weather", getHomeWeather)
	//myRouter.HandleFunc("/article", createNewArticle).Methods("POST")
	//// add our new DELETE endpoint here
	//myRouter.HandleFunc("/article/{id}", deleteArticle).Methods("DELETE")
	//myRouter.HandleFunc("/article/{id}", returnSingleArticle)
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}
func main() {
	services.CreateWeatherRetrievalScheduler()
	handleRequests()
}
