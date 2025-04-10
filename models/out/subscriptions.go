package out

import (
	"XNetVPN-Back/models/db"
)

type Subscriptions struct {
	Subscriptions []subscription `json:"subscriptions"`
}

type subscription struct {
	Id                string   `json:"id" bson:"_id"`
	Name              string   `json:"name" bson:"name"`
	MonthPrice        float64  `json:"month_price" bson:"month_price"`
	AnnuallyPrice     float64  `json:"annually_price" bson:"annually_price"`
	AnnuallyPriceYear float64  `json:"annually_price_year" bson:"annually_price_year"`
	Devices           int      `json:"devices" bson:"devices"`
	Benefits          []string `json:"benefits" bson:"benefits"`
}

func (l *Subscriptions) FillWith(dbSubscriptions []db.Subscription) {
	l.Subscriptions = []subscription{}
	for _, dbSubscription := range dbSubscriptions {
		subscriptionObject := subscription{
			Id:                dbSubscription.Id.Hex(),
			Name:              dbSubscription.Name,
			MonthPrice:        dbSubscription.MonthPrice,
			AnnuallyPrice:     dbSubscription.AnnuallyPrice,
			AnnuallyPriceYear: dbSubscription.AnnuallyPriceYear,
			Devices:           dbSubscription.Devices,
			Benefits:          dbSubscription.Benefits,
		}
		l.Subscriptions = append(l.Subscriptions, subscriptionObject)
	}
}
