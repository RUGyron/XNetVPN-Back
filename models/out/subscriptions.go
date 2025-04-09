package out

import (
	"XNetVPN-Back/models/db"
)

type Subscriptions struct {
	Subscriptions []subscription `json:"subscriptions"`
}

type subscription struct {
	Id       string   `json:"id" bson:"_id"`
	Name     string   `json:"name" bson:"name"`
	Price    float64  `json:"price" bson:"price"`
	Devices  int      `json:"devices" bson:"devices"`
	Benefits []string `json:"benefits" bson:"benefits"`
}

func (l *Subscriptions) FillWith(dbSubscriptions []db.Subscription) {
	l.Subscriptions = []subscription{}
	for _, dbSubscription := range dbSubscriptions {
		subscriptionObject := subscription{
			Id:       dbSubscription.Id.Hex(),
			Name:     dbSubscription.Name,
			Price:    dbSubscription.Price,
			Devices:  dbSubscription.Devices,
			Benefits: dbSubscription.Benefits,
		}
		l.Subscriptions = append(l.Subscriptions, subscriptionObject)
	}
}
