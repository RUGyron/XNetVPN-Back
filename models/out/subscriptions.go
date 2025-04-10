package out

import (
	"XNetVPN-Back/models/db"
)

type Subscriptions struct {
	Subscriptions []subscription `json:"subscriptions"`
}

type subscription struct {
	Id               string   `json:"id" bson:"_id"`
	Name             string   `json:"name" bson:"name"`
	MonthPrice       float64  `json:"monthPrice" bson:"monthPrice"`
	AnnualyPrice     float64  `json:"annualyPrice" bson:"annualyPrice"`
	AnnualyPriceYear float64  `json:"annualyPriceYear" bson:"annualyPriceYear"`
	Devices          int      `json:"devices" bson:"devices"`
	Benefits         []string `json:"benefits" bson:"benefits"`
}

func (l *Subscriptions) FillWith(dbSubscriptions []db.Subscription) {
	l.Subscriptions = []subscription{}
	for _, dbSubscription := range dbSubscriptions {
		subscriptionObject := subscription{
			Id:               dbSubscription.Id.Hex(),
			Name:             dbSubscription.Name,
			MonthPrice:       dbSubscription.MonthPrice,
			AnnualyPrice:     dbSubscription.AnnualyPrice,
			AnnualyPriceYear: dbSubscription.AnnualyPriceYear,
			Devices:          dbSubscription.Devices,
			Benefits:         dbSubscription.Benefits,
		}
		l.Subscriptions = append(l.Subscriptions, subscriptionObject)
	}
}
