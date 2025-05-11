package yookassa

import (
	"XNetVPN-Back/config"
	"XNetVPN-Back/models"
	"XNetVPN-Back/models/in"
	"XNetVPN-Back/repositories/configs"
	"XNetVPN-Back/services/utils/generics"
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"math"
	"net/http"
	"time"
)

type customer struct {
	Email string `json:"email"`
}

type confirmation struct {
	Type      string `json:"type"`
	ReturnURL string `json:"return_url"`
}

type item struct {
	Description string            `json:"description"`
	Amount      in.YookassaAmount `json:"amount"`
	VatCode     int               `json:"vat_code"`
	Quantity    int               `json:"quantity"`
}

type receipt struct {
	Items    []item   `json:"items"`
	Email    string   `json:"email"`
	Customer customer `json:"customer"`
}

type payload struct {
	Amount            in.YookassaAmount `json:"amount"`
	Confirmation      confirmation      `json:"confirmation"`
	Capture           bool              `json:"capture"`
	Description       string            `json:"description"`
	SavePaymentMethod bool              `json:"save_payment_method"`
	Receipt           receipt           `json:"receipt"`
	Metadata          in.Metadata       `json:"metadata"`
}

type response struct {
	Id           string `json:"id"`
	Confirmation struct {
		ConfirmationUrl string `json:"confirmation_url"`
	} `json:"confirmation"`
}

func RequestBillingSave(email string) (*string, *string, error) {
	var paymentResponse response

	amount := 1.0
	amountWithCommission, err := calculateAmountWithCommission(amount)
	if err != nil || amountWithCommission == nil {
		return nil, nil, err
	}

	// payload
	paymentPayload := payload{
		Amount:            in.YookassaAmount{Value: fmt.Sprintf("%.2f", *amountWithCommission), Currency: "RUB"},
		Confirmation:      confirmation{Type: "redirect", ReturnURL: config.Config.SiteUrl},
		Capture:           true,
		Description:       fmt.Sprintf("Привязка счёта для %s", email),
		SavePaymentMethod: true,
		Receipt: receipt{
			Items: []item{{
				Description: "Привязка счёта",
				Amount:      in.YookassaAmount{Value: fmt.Sprintf("%.2f", amount), Currency: "RUB"},
				VatCode:     1,
				Quantity:    1,
			}},
			Email:    email,
			Customer: customer{Email: email},
		},
		Metadata: in.Metadata{Event: models.YKEventType.Save},
	}
	bodyBytes, err := json.Marshal(paymentPayload)
	if err != nil {
		return nil, nil, err
	}

	// headers
	req, err := http.NewRequest("POST", "https://api.yookassa.ru/v3/payments", bytes.NewBuffer(bodyBytes))
	if err != nil {
		return nil, nil, err
	}
	req.Header.Set("Idempotence-Key", uuid.New().String())
	req.Header.Set("Content-Type", "application/json")
	auth := base64.StdEncoding.EncodeToString([]byte(config.Config.YooKassaShopId + ":" + config.Config.YooKassaApiKey))
	req.Header.Set("Authorization", "Basic "+auth)

	// do request
	client := &http.Client{Timeout: time.Duration(config.Config.TimeoutExternalHttp) * time.Millisecond}
	resp, err := client.Do(req)
	if err != nil || resp == nil {
		return nil, nil, err
	}

	// bind response
	err = generics.BindStructWithResponse(*resp, &paymentResponse, true)
	if err != nil {
		return nil, nil, err
	}

	return &paymentResponse.Id, &paymentResponse.Confirmation.ConfirmationUrl, nil
}

// calculateAmountWithCommission 100 + 4% = 104.17
func calculateAmountWithCommission(targetAmount float64) (*float64, error) {
	// yookassa rate
	yookassaRate, err := configs.FindYookassaRate()
	if err != nil || yookassaRate == nil {
		return nil, err
	}

	result := targetAmount / ((100 - *yookassaRate) / 100)
	result = math.Ceil(result*100) / 100
	return &result, err
}
