package out

import (
	"XNetVPN-Back/models/db"
)

type Config struct {
	Id string `bson:"_id" json:"id"`

	PrivateKey string `bson:"private_key" json:"private_key"`
	Address    string `bson:"address" json:"address"`
	DNS        string `bson:"dns" json:"dns"`

	Endpoint     string   `bson:"endpoint" json:"endpoint"`
	AllowedIPs   []string `bson:"allowed_ips" json:"allowed_ips"`
	PresharedKey string   `bson:"preshared_key" json:"preshared_key"`
	Publickey    string   `bson:"public_key" json:"public_key"`
}

func (c *Config) FillWith(config db.Config) {
	c.Id = config.Id.Hex()
	c.PrivateKey = config.PrivateKey
	c.Address = config.Address
	c.DNS = config.DNS
	c.Endpoint = config.Endpoint
	c.AllowedIPs = config.AllowedIPs
	c.PresharedKey = config.PresharedKey
	c.Publickey = config.Publickey
}
