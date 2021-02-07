package payload

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"group-management-api/domain/model"
)

// swagger:model UnregisterUserPayload
type UnregisterUserPayload struct {
	// The current users password.
	//
	// required: true
	// minimum: 4
	// maximum: 120
	// example: password
	Password string `json:"password"`
}

func (uup UnregisterUserPayload) Validate() error {
	return validation.ValidateStruct(&uup,
		validation.Field(&uup.Password, model.UserPasswordRule...),
	)
}
