package out

import (
	"time"
	"xnet-vpn/models"
	"xnet-vpn/models/db"
)

type Login struct {
	User struct {
		Id             string    `bson:"_id" json:"id"`
		SubscriptionId *string   `bson:"subscription_id" json:"subscription_id"`
		CreatedAt      time.Time `bson:"created_at" json:"created_at"`
	}
	Tokens models.Tokens `json:"tokens" bson:"tokens"`
}

func (l *Login) FillWith(user *db.User) {
	l.User.Id = user.Id.Hex()
	l.User.CreatedAt = user.CreatedAt
	if user.SubscriptionId != nil {
		subId := user.SubscriptionId.Hex()
		l.User.SubscriptionId = &subId
	}
}
