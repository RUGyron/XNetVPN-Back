package out

import (
	"XNetVPN-Back/models/db"
	"time"
)

type Profile struct {
	Id              string           `json:"_id"`
	CreatedAt       time.Time        `json:"created_at"`
	SubscriptionEnd *time.Time       `json:"subscription_end"`
	Subscription    userSubscription `json:"subscription"`
	Devices         []device         `json:"devices"`
}

type device struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type userSubscription struct {
	Name    string  `json:"name"`
	Price   float64 `json:"price"`
	Devices int     `json:"devices"`
}

func (p *Profile) FillWith(user *db.User, devices []db.Device, subscription db.Subscription) {
	p.Id = user.Id.Hex()
	p.CreatedAt = user.CreatedAt
	p.SubscriptionEnd = user.SubscriptionEnd
	p.Devices = make([]device, 0, len(devices))
	for _, d := range devices {
		p.Devices = append(p.Devices, device{
			Name: d.Name,
			Type: d.Type,
		})
	}
	p.Subscription.Name = subscription.Name
	p.Subscription.Devices = subscription.Devices
	p.Subscription.Price = subscription.Price
}
