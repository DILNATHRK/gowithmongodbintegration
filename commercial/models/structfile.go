package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type City struct {
	CityID   primitive.ObjectID `bson:"_id"`
	CityName string             `bson:"city_name"`
}


