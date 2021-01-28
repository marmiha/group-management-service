package payload

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"group-management-api/domain/model"
)

type UnregisterUserPayload struct {
	Password string `json:"password"`
}

func (uup UnregisterUserPayload) Validate() error {
	return validation.ValidateStruct(&uup,
		validation.Field(&uup.Password, model.UserPasswordRule...),
	)
}
