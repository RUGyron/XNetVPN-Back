package repo_subscriptions

import (
	"XNetVPN-Back/config"
	"XNetVPN-Back/models/db"
	"XNetVPN-Back/repositories"
	"XNetVPN-Back/services/utils/generics"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func FindSubscriptions() ([]db.Subscription, error) {
	collection := repositories.MajorityClient.Database(config.Config.MongoDatabase).Collection(config.Config.MongoCollectionSubscriptions)
	filter := bson.M{}

	opts := options.Find().SetMaxTime(time.Duration(config.Config.TimeoutMongoQueryInside) * time.Millisecond)
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(config.Config.TimeoutMongoQuery)*time.Millisecond)
	defer cancel()

	var subscriptions []db.Subscription
	cursor, err := collection.Find(ctx, filter, opts)
	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		return make([]db.Subscription, 0), err
	}

	if err = generics.BindStructArrayWithCursor(ctx, cursor, &subscriptions); err != nil {
		return make([]db.Subscription, 0), err
	}

	return subscriptions, nil
}
