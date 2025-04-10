package in

import "XNetVPN-Back/services/utils"

type Login struct {
	Key *string `json:"key"`
}

func (payload Login) Validate() (bool, []string) {
	return utils.ValidateStruct(payload)
}
