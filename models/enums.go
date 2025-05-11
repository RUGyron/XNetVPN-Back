package models

var RegexRules = map[string]string{}

type ProductSwitchResult struct {
	AmountToPay float64
	NewCredit   float64
	ApplyNow    bool
}

var YKEventType = struct {
	Save   string
	Pay    string
	Delete string
}{
	Save:   "save",
	Pay:    "pay",
	Delete: "delete",
}

var YKEventStatus = struct {
	Pending string
	Success string
	Error   string
}{
	Pending: "pending",
	Success: "success",
	Error:   "error",
}
