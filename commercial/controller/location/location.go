package location

import (
	"commercial-propfloor/database"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"

	// "log"
	"github.com/gin-gonic/gin"
)

type CityName struct {
	City_name string `json:"city_name"`
}

func CreateCity() gin.HandlerFunc {
	return func(c *gin.Context) {
		var City_name = "Ankit"
		fmt.Println(City_name)

		// _, err := database.Mongomain()
		// if err != nil{
		// 	panic(err)
		// }
		client, cc, cancel, err := database.Mongoconnect("mongodb://localhost:27017")

		if err != nil {
			return
		}

		collection := client.Database("city-name").Collection("city")

		res, err := collection.InsertOne(context.Background(), bson.M{"city_name": "Pune"})
		if err != nil {
			return err
		}
		
		id := res.InsertedID

		fmt.Println(id)

		database.Mongoclose(client, cc, cancel)
		return
	}
}
