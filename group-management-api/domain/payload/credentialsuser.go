package payload

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"group-management-api/domain/model"
)

type CredentialsUserPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (cup CredentialsUserPayload) Validate() error {
	return validation.ValidateStruct(&cup,
		validation.Field(&cup.Email, model.UserEmailRequiredRule...),
		validation.Field(&cup.Password, model.UserPasswordRule...),
	)
}
