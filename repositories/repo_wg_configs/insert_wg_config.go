package repo_wg_configs

import (
	"XNetVPN-Back/config"
	"XNetVPN-Back/models/in"
	"XNetVPN-Back/repositories"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func InsertWgConfig(payload *in.ConfigResponse) (primitive.ObjectID, error) {
	collection := repositories.MajorityClient.
		Database(config.Config.MongoDatabase).
		Collection(config.Config.MongoCollectionWgConfigs)

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(config.Config.TimeoutMongoQuery)*time.Millisecond)
	defer cancel()

	result, err := collection.InsertOne(ctx, bson.M{
		"identifier":           payload.Identifier,
		"interface_identifier": payload.InterfaceIdentifier,
		"private_key":          payload.PrivateKey,
		"address":              payload.Addresses[0],
		"dns":                  payload.Dns.Value[0],
		"endpoint":             payload.Endpoint.Value,
		"allowed_ips":          payload.AllowedIPs.Value,
		"preshared_key":        payload.PresharedKey,
		"public_key":           payload.EndpointPublicKey.Value,
	})
	if err != nil {
		return primitive.NilObjectID, err
	}

	insertedID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return primitive.NilObjectID, errors.New("failed to assert inserted ID as ObjectID")
	}

	return insertedID, nil
}
