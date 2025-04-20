package db

import "go.mongodb.org/mongo-driver/bson/primitive"

type Subscription struct {
	Id                primitive.ObjectID `json:"id" bson:"_id"`
	Name              string             `json:"name" bson:"name"`
	MonthPrice        float64            `json:"month_price" bson:"month_price"`
	AnnuallyPrice     float64            `json:"annually_price" bson:"annually_price"`
	AnnuallyPriceYear float64            `json:"annually_price_year" bson:"annually_price_year"`
	Devices           int                `json:"devices" bson:"devices"`
	Benefits          []string           `json:"benefits" bson:"benefits"`
}

type Product struct {
	Id     primitive.ObjectID `json:"id" bson:"_id"`
	Name   string             `json:"name" bson:"name"`
	Price  float64            `json:"price" bson:"price"`
	Annual bool               `json:"annual" bson:"annual"`
	Rank   int                `json:"rank" bson:"rank"`
}
