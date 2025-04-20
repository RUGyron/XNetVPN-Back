package models

var RegexRules = map[string]string{}

type ProductSwitchResult struct {
	AmountToPay float64
	NewCredit   float64
	ApplyNow    bool
}
