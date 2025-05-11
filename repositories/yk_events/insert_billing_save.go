package yk_events

import (
	"XNetVPN-Back/config"
	"XNetVPN-Back/models"
	"XNetVPN-Back/repositories"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

func InsertBillingSave(ykId, email string) error {
	collection := repositories.MajorityClient.Database(config.Config.MongoDatabase).Collection(config.Config.MongoCollectionYKEvents)
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(config.Config.TimeoutMongoQuery)*time.Millisecond)
	defer cancel()
	_, err := collection.InsertOne(ctx, bson.M{
		"yk_id":        ykId,
		"event":        models.YKEventType.Save,
		"status":       models.YKEventStatus.Pending,
		"created_at":   time.Now().UTC(),
		"rub_amount":   1.0,
		"billing_type": nil,
		"billing_id":   nil,
		"email":        email,
	})
	return err
}
