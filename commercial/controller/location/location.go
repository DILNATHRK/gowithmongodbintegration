package location

import (
	"commercial-propfloor/database"
	"commercial-propfloor/models"
	"context"
	"fmt"
	"log"
	"reflect"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

// var Client *mongo.Client = CreateCity()

// func InsertInDatabase() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		client, _, _, _ := database.Mongoconnect("mongodb://localhost:27017")
// 		collection := client.Database("city-name").Collection("city")
// 		res, errr := collection.InsertOne(context.Background(), bson.M{"city_name": "Trivandrum"})
// 		if errr != nil {
// 			log.Fatal(errr)
// 		}
// 		id := res.InsertedID
// 		fmt.Println("inserted-cityid  ", id)
// 		return
// 	}
// }

func InsertInDatabase() gin.HandlerFunc {
	return func(c *gin.Context) {

		client, _, _, _ := database.Mongoconnect("mongodb://localhost:27017")
		coll := client.Database("building_details").Collection("commertial_building")
		// filter := bson.D{{"city_name", "Pune"}}
		filter := bson.D{}

		var city []models.City
		cursor, err := coll.Find(context.Background(), filter)
		cursor.All(context.Background(), &city)

		if err != nil {
			if err == mongo.ErrNoDocuments {
				fmt.Println("Ankit")
				// This error means your query did not match any documents.
				return
			}
			panic(err)
		}
		fmt.Println("datatype", reflect.TypeOf(city))
		fmt.Println(city)
		fmt.Println("Selected")
		fmt.Println(city[0].CityName)
		return
	}
}

func DeleteInDatabase() gin.HandlerFunc {
	return func(c *gin.Context) {
		client, _, _, _ := database.Mongoconnect("mongodb://localhost:27017")
		collection := client.Database("city-name").Collection("city")
		res, errr := collection.DeleteOne(context.Background(), bson.M{"city_name": "trivandrum"})
		if errr != nil {
			log.Fatal(errr)
		}

		fmt.Println(res)

		return
	}
}

func UpdateInDatabase() gin.HandlerFunc {
	return func(c *gin.Context) {
		client, _, _, _ := database.Mongoconnect("mongodb://localhost:27017")
		collection := client.Database("city-name").Collection("city")

		id, _ := primitive.ObjectIDFromHex("641fd3ad838407c14c716e4d")
		filter := bson.D{{"_id", id}}
		update := bson.D{{"$set", bson.D{{"city_name", "Dubai"}}}}

		res, errr := collection.UpdateOne(context.Background(), filter, update)
		if errr != nil {
			log.Fatal(errr)
		}

		fmt.Println(res)

		return
	}
}

func Selectindatabase() gin.HandlerFunc {
	return func(c *gin.Context) {
		client, ctx, cancel, _ := database.Mongoconnect("mongodb://localhost:27017")

		collection := client.Database("city-name").Collection("city")
		cursor, err := collection.Find(ctx, bson.M{})
		if err != nil {
			log.Println(err)
		}

		var city []bson.M
		if err = cursor.All(ctx, &city); err != nil {
			log.Println(err)
		}
		fmt.Println()
		fmt.Println("datatype", reflect.TypeOf(city))
		fmt.Println("cities\n", city)
		fmt.Println("****************************************************")
		fmt.Printf("selected city-----\n", city[0]["city_name"])
		fmt.Println("****************************************************")
		database.Mongoclose(client, ctx, cancel)
		return
	}

}

func Queryindatabase() gin.HandlerFunc {
	return func(c *gin.Context) {
		client, ctx, cancel, errr := database.Mongoconnect("mongodb://localhost:27017")
		if errr != nil {
			log.Panic(errr)
		}

		//var ctx, _ = context.WithTimeout(context.Background(), 20*time.Second)
		collection := client.Database("city-name").Collection("city")
		var city []bson.M
		filter, err := collection.Find(ctx, bson.M{"city_name": "Bangalore"})
		if err != nil {
			log.Println(err)
		}
		if err = filter.All(ctx, &city); err != nil {
			log.Println(err)
		}
		fmt.Println("querried document ", city)
		fmt.Println()

		database.Mongoclose(client, ctx, cancel)

		return
	}

}
