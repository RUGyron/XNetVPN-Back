package in

import "XNetVPN-Back/services/utils"

type AddDevice struct {
	Name       string `json:"name"`
	Type       string `json:"type"`
	Identifier string `json:"identifier"`
}

func (payload AddDevice) Validate() (bool, []string) {
	return utils.ValidateStruct(payload)
}
