package db

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	Id                    primitive.ObjectID  `bson:"_id"`
	SubscriptionId        *primitive.ObjectID `bson:"subscription_id"`
	CreatedAt             time.Time           `bson:"created_at"`
	SubscriptionExpiresAt *time.Time          `bson:"subscription_expires_at"`
}
