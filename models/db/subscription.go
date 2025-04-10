package db

import "go.mongodb.org/mongo-driver/bson/primitive"

type Subscription struct {
	Id               primitive.ObjectID `json:"id" bson:"_id"`
	Name             string             `json:"name" bson:"name"`
	MonthPrice       float64            `json:"monthPrice" bson:"monthPrice"`
	AnnualyPrice     float64            `json:"annualyPrice" bson:"annualyPrice"`
	AnnualyPriceYear float64            `json:"annualyPriceYear" bson:"annualyPriceYear"`
	Devices          int                `json:"devices" bson:"devices"`
	Benefits         []string           `json:"benefits" bson:"benefits"`
}
