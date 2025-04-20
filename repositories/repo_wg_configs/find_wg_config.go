package repo_wg_configs

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

func FindWgConfigById(configId *primitive.ObjectID) (*db.WgConfig, error) {
	collection := repositories.MajorityClient.Database(config.Config.MongoDatabase).Collection(config.Config.MongoCollectionWgConfigs)
	filter := bson.M{"_id": configId}
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(config.Config.TimeoutMongoQuery)*time.Millisecond)
	defer cancel()
	var result *db.WgConfig
	err := collection.FindOne(ctx, filter, options.FindOne().SetMaxTime(time.Duration(config.Config.TimeoutMongoQueryInside)*time.Millisecond)).Decode(&result)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}
	return result, nil
}
