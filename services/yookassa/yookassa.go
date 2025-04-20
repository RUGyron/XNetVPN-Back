package yookassapackage

import (
	"XNetVPN-Back/config"
	"XNetVPN-Back/services/utils/generics"
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"net/http"
	"time"
)

type amount struct {
	Value    string `json:"value"`
	Currency string `json:"currency"`
}

type customer struct {
	Email string `json:"email"`
}

type confirmation struct {
	Type      string `json:"type"`
	ReturnURL string `json:"return_url"`
}

type item struct {
	Description string `json:"description"`
	Amount      amount `json:"amount"`
	VatCode     int    `json:"vat_code"`
	Quantity    int    `json:"quantity"`
}

type receipt struct {
	Items    []item   `json:"items"`
	Email    string   `json:"email"`
	Customer customer `json:"customer"`
}

type payload struct {
	Amount            amount       `json:"amount"`
	Confirmation      confirmation `json:"confirmation"`
	Capture           bool         `json:"capture"`
	Description       string       `json:"description"`
	SavePaymentMethod bool         `json:"save_payment_method"`
	Receipt           receipt      `json:"receipt"`
}

type response struct {
}

func RequestBillingSave(email string) error {
	var paymentResponse response
	// payload
	paymentPayload := payload{
		Amount:            amount{Value: "1.00", Currency: "RUB"},
		Confirmation:      confirmation{Type: "redirect", ReturnURL: config.Config.SiteUrl},
		Capture:           true,
		Description:       fmt.Sprintf("Привязка счёта для %s", email),
		SavePaymentMethod: true,
		Receipt: receipt{
			Items: []item{{
				Description: "Привязка счёта",
				Amount:      amount{Value: "1.00", Currency: "RUB"},
				VatCode:     1,
				Quantity:    1,
			}},
			Email:    email,
			Customer: customer{Email: email},
		},
	}
	bodyBytes, err := json.Marshal(paymentPayload)
	if err != nil {
		return err
	}

	// headers
	req, err := http.NewRequest("POST", "https://api.yookassa.ru/v3/payments", bytes.NewBuffer(bodyBytes))
	if err != nil {
		return err
	}
	req.Header.Set("Idempotence-Key", uuid.New().String())
	req.Header.Set("Content-Type", "application/json")
	auth := base64.StdEncoding.EncodeToString([]byte(config.Config.YooKassaShopId + ":" + config.Config.YooKassaApiKey))
	req.Header.Set("Authorization", "Basic "+auth)

	// do request
	client := &http.Client{Timeout: time.Duration(config.Config.TimeoutExternalHttp) * time.Millisecond}
	resp, err := client.Do(req)
	if err != nil || resp == nil {
		return err
	}

	// bind response
	err = generics.BindStructWithResponse(*resp, &paymentResponse, true)
	if err != nil {
		return err
	}

	return nil
}
