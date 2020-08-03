package services

import (
	"api-test/api"
	"api-test/dal"
	"fmt"
 "github.com/robfig/cron/v3"
)

func UpdateWeather(){
	res,err := api.GetWeather("Aarhus")
		if err != nil {
			fmt.Printf("Invalid response from Weather Service")
			return
		}
	result:= dal.AddWeatherEntry(*res)
	if result != nil{
		fmt.Println("Successfully added new weather entry")
	}
}

func CreateWeatherRetrievalScheduler() {
	c:= cron.New()
	id,err :=c.AddFunc("* * * * *", func() {
		UpdateWeather()
	})
	c.Start()
	if err !=nil {
		fmt.Println(err)
	} else {
		fmt.Println("Successfully created cron job with id "+string(id))
	}
}
