package repo_devices

import (
	"XNetVPN-Back/config"
	"XNetVPN-Back/models/in"
	"XNetVPN-Back/repositories"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func InsertDevice(payload in.AddDevice, userId primitive.ObjectID) error {
	collection := repositories.MajorityClient.Database(config.Config.MongoDatabase).Collection(config.Config.MongoCollectionDevices)
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(config.Config.TimeoutMongoQuery)*time.Millisecond)
	defer cancel()

	tgUser := bson.M{
		"user_id":    userId,
		"name":       payload.Name,
		"type":       payload.Type,
		"identifier": payload.Identifier,
	}

	_, err := collection.InsertOne(ctx, tgUser)
	if err != nil {
		return err
	}

	return nil
}
