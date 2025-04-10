package db

import "go.mongodb.org/mongo-driver/bson/primitive"

type Subscription struct {
	Id                primitive.ObjectID `json:"id" bson:"_id"`
	Name              string             `json:"name" bson:"name"`
	MonthPrice        float64            `json:"monthPrice" bson:"monthPrice"`
	AnnuallyPrice     float64            `json:"annuallyPrice" bson:"annuallyPrice"`
	AnnuallyPriceYear float64            `json:"annuallyPriceYear" bson:"annuallyPriceYear"`
	Devices           int                `json:"devices" bson:"devices"`
	Benefits          []string           `json:"benefits" bson:"benefits"`
}
