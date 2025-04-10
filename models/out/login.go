package out

import (
	"XNetVPN-Back/models"
)

type Login struct {
	User   Profile       `bson:"user" json:"user"`
	Tokens models.Tokens `bson:"tokens" json:"tokens"`
}
