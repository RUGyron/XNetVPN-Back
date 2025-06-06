package out

import (
	"XNetVPN-Back/models/db"
	"time"
)

type Profile struct {
	Id                    string            `json:"id" bson:"_id"`
	CreatedAt             time.Time         `json:"created_at"`
	SubscriptionExpiresAt *time.Time        `json:"subscription_expires_at"`
	Subscription          *userSubscription `json:"subscription"`
	Devices               []device          `json:"devices"`
}

type device struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type userSubscription struct {
	Name              string  `json:"name"`
	MonthPrice        float64 `json:"month_price"`
	AnnuallyPrice     float64 `json:"annually_price"`
	AnnuallyPriceYear float64 `json:"annually_price_year"`
	Devices           int     `json:"devices"`
}

func (p *Profile) FillWith(user *db.User, devices []db.Device, subscription *db.Subscription) {
	p.Id = user.Id.Hex()
	p.CreatedAt = user.CreatedAt
	p.SubscriptionExpiresAt = user.SubscriptionExpiresAt
	p.Devices = make([]device, 0, len(devices))
	for _, d := range devices {
		p.Devices = append(p.Devices, device{
			Name: d.Name,
			Type: d.Type,
		})
	}
	if subscription != nil {
		p.Subscription = &userSubscription{
			Name:              subscription.Name,
			MonthPrice:        subscription.MonthPrice,
			AnnuallyPrice:     subscription.AnnuallyPrice,
			AnnuallyPriceYear: subscription.AnnuallyPriceYear,
			Devices:           subscription.Devices,
		}
	}
}
