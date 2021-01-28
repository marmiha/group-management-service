package payload

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"group-management-api/domain/model"
)

type ChangePasswordPayload struct {
	CurrentPassword string `json:"current_password"`
	NewPassword     string `json:"new_password"`
}

func (cpp ChangePasswordPayload) Validate() error {
	return validation.ValidateStruct(&cpp,
		validation.Field(&cpp.NewPassword, model.UserPasswordRule...),
		validation.Field(&cpp.CurrentPassword, model.UserPasswordRule...),
	)
}
