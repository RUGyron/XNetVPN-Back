package db

import "go.mongodb.org/mongo-driver/bson/primitive"

type WgConfig struct {
	Id                  primitive.ObjectID `bson:"_id" json:"id"`
	Identifier          string             `bson:"identifier" json:"identifier"`
	InterfaceIdentifier string             `bson:"interface_identifier" json:"interface_identifier"`

	PrivateKey string `bson:"private_key" json:"private_key"`
	Address    string `bson:"address" json:"address"`
	DNS        string `bson:"dns" json:"dns"`

	Endpoint     string   `bson:"endpoint" json:"endpoint"`
	AllowedIPs   []string `bson:"allowed_ips" json:"allowed_ips"`
	PresharedKey string   `bson:"preshared_key" json:"preshared_key"`
	Publickey    string   `bson:"public_key" json:"public_key"`
}
