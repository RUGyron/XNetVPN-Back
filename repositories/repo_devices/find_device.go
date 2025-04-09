package repo_devices

import (
	"XNetVPN-Back/config"
	"XNetVPN-Back/models/db"
	"XNetVPN-Back/repositories"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func FindDevice(identifier string) (*db.Device, error) {
	collection := repositories.MajorityClient.Database(config.Config.MongoDatabase).Collection(config.Config.MongoCollectionSubscriptions)
	filter := bson.M{"identifier": identifier}
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(config.Config.TimeoutMongoQuery)*time.Millisecond)
	defer cancel()
	var device *db.Device
	err := collection.FindOne(ctx, filter, options.FindOne().SetMaxTime(time.Duration(config.Config.TimeoutMongoQueryInside)*time.Millisecond)).Decode(&device)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}
	return device, nil
}
