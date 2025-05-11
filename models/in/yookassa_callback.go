package in

type YookassaCallback struct {
	Event  string       `json:"event"`
	Object callbackData `json:"object"`
}

type recipient struct {
	AccountId string `json:"account_id"`
	GatewayId string `json:"gateway_id"`
}

type paymentMethod struct {
	Id            string `json:"id"`
	Type          string `json:"type"`
	Saved         bool   `json:"saved"`
	Status        string `json:"status"`
	Title         string `json:"title"`
	AccountNumber string `json:"account_number"`
}

type callbackData struct {
	Id            string         `json:"id"`
	Status        string         `json:"status"`
	IncomeAmount  YookassaAmount `json:"income_amount"`
	Description   string         `json:"description"`
	Recipient     recipient      `json:"recipient"`
	PaymentMethod paymentMethod  `json:"payment_method"`
	Paid          bool           `json:"paid"`
	Metadata      Metadata       `json:"metadata"`
}

type YookassaAmount struct {
	Value    string `json:"value"`
	Currency string `json:"currency"`
}

type Metadata struct {
	Event string `json:"event"`
}
