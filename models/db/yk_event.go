package db

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type YKEvent struct {
	Id          primitive.ObjectID `bson:"_id,omitempty"`
	YKId        string             `bson:"yk_id"`
	Event       string             `bson:"event"`
	Status      string             `bson:"status"`
	CreatedAt   time.Time          `bson:"created_at"`
	RubAmount   float64            `bson:"rub_amount"`
	BillingType *string            `bson:"billing_type"`
	BillingId   *string            `bson:"billing_id"`
	Email       string             `bson:"email"`
}
