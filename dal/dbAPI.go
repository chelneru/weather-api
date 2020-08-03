package dal

import (
	"api-test/api"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type HomeWeatherObject struct {
	Humidity float64
	Temperature float64
	Time string
}
func AddWeatherEntry(weatherObj api.WeatherObject) interface{} {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		fmt.Printf("Unable to connect to the database")
		return nil
	}
	collection := client.Database("weather").Collection("weather_temps")
	res, err := collection.InsertOne(ctx,weatherObj)
	if err != nil {
		fmt.Printf("Unable to insert entry in the collection")
		return nil
	}
	id := res.InsertedID
	return id
}

func RetrieveLatestWeatherEntry() *api.WeatherObject {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		fmt.Printf("Unable to connect to the database")
		return nil
	}

	opts := options.Find()

	// Sort by `_id` field descending
	opts.SetSort(bson.D{{"_id", -1}})

	// Limit by 10 documents only
	opts.SetLimit(1)
	collection := client.Database("weather").Collection("weather_temps")
	var weatherEntry api.WeatherObject

	cursor, err := collection.Find(context.Background(), bson.M{}, opts)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	for cursor.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		err := cursor.Decode(&weatherEntry)
		if err != nil {
			log.Fatal(err)
		}
	}
	cursor.Close(context.TODO())
	ctx.Done()
	return &weatherEntry

}

func RetrieveLatestHomeWeatherEntry() *HomeWeatherObject {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		fmt.Printf("Unable to connect to the database")
		return nil
	}

	opts := options.Find()

	// Sort by `_id` field descending
	opts.SetSort(bson.D{{"_id", -1}})

	// Limit by 10 documents only
	opts.SetLimit(1)
	collection := client.Database("weather").Collection("home_weather_temps")
	var weatherEntry HomeWeatherObject

	cursor, err := collection.Find(context.Background(), bson.M{}, opts)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	for cursor.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		err := cursor.Decode(&weatherEntry)
		if err != nil {
			log.Fatal(err)
		}
	}
	cursor.Close(context.TODO())
	ctx.Done()
	return &weatherEntry

}