package payload

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"group-management-api/domain/model"
)

type RegisterUserPayload struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

func (rup RegisterUserPayload) Validate() error {
	return validation.ValidateStruct(&rup,
		validation.Field(&rup.Email, model.UserEmailRequiredRule...),
		validation.Field(&rup.Name, model.UserNameRule...),
		validation.Field(&rup.Password, model.UserPasswordRule...),
	)
}
