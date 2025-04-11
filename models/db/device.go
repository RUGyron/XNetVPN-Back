package db

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Device struct {
	Id         primitive.ObjectID  `bson:"_id"`
	UserId     primitive.ObjectID  `bson:"user_id"`
	Name       string              `bson:"name"`
	Type       string              `bson:"type"`
	Identifier string              `bson:"identifier"`
	ConfigId   *primitive.ObjectID `bson:"config_id"`
}
