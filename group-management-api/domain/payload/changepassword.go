package payload

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"group-management-api/domain/model"
)

// swagger:model ChangePasswordPayload
type ChangePasswordPayload struct {
	// The current password of the user.
	//
	// required: true
	// minimum: 4
	// maximum: 120
	// example: password
	CurrentPassword string `json:"current_password"`

	// The new password for the user.
	//
	// required: true
	// minimum: 4
	// maximum: 120
	// example: new_password
	NewPassword string `json:"new_password"`
}

func (cpp ChangePasswordPayload) Validate() error {
	return validation.ValidateStruct(&cpp,
		validation.Field(&cpp.NewPassword, model.UserPasswordRule...),
		validation.Field(&cpp.CurrentPassword, model.UserPasswordRule...),
	)
}
