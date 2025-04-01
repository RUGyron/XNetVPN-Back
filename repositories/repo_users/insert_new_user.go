package repo_users

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
	"xnet-vpn-back/config"
	"xnet-vpn-back/repositories"
)

func InsertNewUser() (*primitive.ObjectID, error) {
	collection := repositories.MajorityClient.Database(config.Config.MongoDatabase).Collection(config.Config.MongoCollectionUsers)
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(config.Config.TimeoutMongoQuery)*time.Millisecond)
	defer cancel()
	result, err := collection.InsertOne(ctx, bson.M{
		"subscription_id": nil,
		"created_at":      time.Now().UTC(),
	})
	if err != nil {
		return nil, err
	}
	insertedId, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, errors.New("failed to get inserted ID")
	}
	return &insertedId, nil
}
