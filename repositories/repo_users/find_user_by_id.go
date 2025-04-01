package repo_users

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
	"xnet-vpn-back/config"
	"xnet-vpn-back/models/db"
	"xnet-vpn-back/repositories"
)

func FindUserById(userId primitive.ObjectID) (*db.User, error) {
	collection := repositories.MajorityClient.Database(config.Config.MongoDatabase).Collection(config.Config.MongoCollectionUsers)
	filter := bson.M{"_id": userId}
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(config.Config.TimeoutMongoQuery)*time.Millisecond)
	defer cancel()
	var result *db.User
	err := collection.FindOne(ctx, filter, options.FindOne().SetMaxTime(time.Duration(config.Config.TimeoutMongoQueryInside)*time.Millisecond)).Decode(&result)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}
	return result, nil
}
