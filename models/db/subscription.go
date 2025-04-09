package db

import "go.mongodb.org/mongo-driver/bson/primitive"

type Subscription struct {
	Id       primitive.ObjectID `json:"id" bson:"_id"`
	Name     string             `json:"name" bson:"name"`
	Price    float64            `json:"price" bson:"price"`
	Devices  int                `json:"devices" bson:"devices"`
	Benefits []string           `json:"benefits" bson:"benefits"`
}
