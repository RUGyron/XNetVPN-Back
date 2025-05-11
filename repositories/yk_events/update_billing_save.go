package yk_events

import (
	"XNetVPN-Back/config"
	"XNetVPN-Back/models"
	"XNetVPN-Back/models/in"
	"XNetVPN-Back/repositories"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

func UpdateBillingSave(payload in.YookassaCallback) error {
	collection := repositories.MajorityClient.Database(config.Config.MongoDatabase).Collection(config.Config.MongoCollectionYKEvents)
	filter := bson.M{"yk_id": payload.Object.Id, "status": models.YKEventStatus.Pending}
	update := bson.M{"$set": bson.M{
		"status":       getLocalStatus(payload.Object.Status),
		"rub_amount":   payload.Object.IncomeAmount.Value,
		"billing_type": payload.Object.PaymentMethod.Type,
		"billing_id":   payload.Object.PaymentMethod.Id,
	}}
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(config.Config.TimeoutMongoQuery)*time.Millisecond)
	defer cancel()
	cursor, err := collection.UpdateOne(ctx, filter, update)
	if err != nil || cursor == nil || cursor.ModifiedCount == 0 {
		return errors.New("failed to update yk status")
	}
	return nil
}

func getLocalStatus(ykStatus string) string {
	switch ykStatus {
	case "succeeded":
		return models.YKEventStatus.Success
	case "canceled":
		return models.YKEventStatus.Error
	default:
		return models.YKEventStatus.Pending
	}
}
