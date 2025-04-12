package repo_devices

import (
	"XNetVPN-Back/config"
	"XNetVPN-Back/repositories"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func UpdateDeviceConfig(idDevice, idConfig primitive.ObjectID) error {
	collection := repositories.MajorityClient.Database(config.Config.MongoDatabase).Collection(config.Config.MongoCollectionDevices)
	filter := bson.M{"_id": idDevice, "config_id": nil}
	update := bson.M{"$set": bson.M{"config_id": idConfig}}
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(config.Config.TimeoutMongoQuery)*time.Millisecond)
	defer cancel()
	_, err := collection.UpdateOne(ctx, filter, update)
	return err
}
