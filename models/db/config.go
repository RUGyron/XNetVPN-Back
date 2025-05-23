package db

import "go.mongodb.org/mongo-driver/bson/primitive"

type FloatConfig struct {
	Id    primitive.ObjectID `bson:"_id,omitempty"`
	Key   string             `bson:"key"`
	Value float64            `bson:"value"`
}
