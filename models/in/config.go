package in

import (
	"XNetVPN-Back/services/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Config struct {
	DeviceId primitive.ObjectID `json:"device_id"`
}

func (payload Config) Validate() (bool, []string) {
	return utils.ValidateStruct(payload)
}
