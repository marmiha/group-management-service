package payload

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"group-management-api/domain/model"
)

// swagger:model RegisterUserPayload
type RegisterUserPayload struct {

	// Email used to login the user.
	//
	// required: true
	// example: dwight.schrute@dunder-mifflin.com
	Email string `json:"email"`

	// User name.
	//
	// minimum: 3
	// maximum: 40
	// example: Dwight Schrute
	Name string `json:"name"`

	// Users passwords.
	//
	// required: true
	// minimum: 4
	// maximum: 120
	// example: password
	Password string `json:"password"`
}

func (rup RegisterUserPayload) Validate() error {
	return validation.ValidateStruct(&rup,
		validation.Field(&rup.Email, model.UserEmailRequiredRule...),
		validation.Field(&rup.Name, model.UserNameRule...),
		validation.Field(&rup.Password, model.UserPasswordRule...),
	)
}
