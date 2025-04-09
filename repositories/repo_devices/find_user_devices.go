package repo_devices

import (
	"XNetVPN-Back/config"
	"XNetVPN-Back/models/db"
	"XNetVPN-Back/repositories"
	"XNetVPN-Back/services/utils/generics"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func FindUserDevices(userId primitive.ObjectID) ([]db.Device, error) {
	collection := repositories.MajorityClient.Database(config.Config.MongoDatabase).Collection(config.Config.MongoCollectionSubscriptions)
	filter := bson.M{"user_id": userId}

	opts := options.Find().SetMaxTime(time.Duration(config.Config.TimeoutMongoQueryInside) * time.Millisecond)
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(config.Config.TimeoutMongoQuery)*time.Millisecond)
	defer cancel()

	var devices []db.Device
	cursor, err := collection.Find(ctx, filter, opts)
	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		return make([]db.Device, 0), err
	}

	if err = generics.BindStructArrayWithCursor(ctx, cursor, &devices); err != nil {
		return make([]db.Device, 0), err
	}

	return devices, nil
}
