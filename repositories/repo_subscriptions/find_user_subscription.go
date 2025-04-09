package repo_subscriptions

import (
	"XNetVPN-Back/config"
	"XNetVPN-Back/models/db"
	"XNetVPN-Back/repositories"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func FindUserSubscription(id primitive.ObjectID) (*db.Subscription, error) {
	collection := repositories.MajorityClient.Database(config.Config.MongoDatabase).Collection(config.Config.MongoCollectionSubscriptions)
	filter := bson.M{"_id": id}
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(config.Config.TimeoutMongoQuery)*time.Millisecond)
	defer cancel()
	var subscription *db.Subscription
	err := collection.FindOne(ctx, filter, options.FindOne().SetMaxTime(time.Duration(config.Config.TimeoutMongoQueryInside)*time.Millisecond)).Decode(&subscription)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}
	return subscription, nil
}
