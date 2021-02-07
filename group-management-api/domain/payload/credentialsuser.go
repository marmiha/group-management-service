package payload

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"group-management-api/domain/model"
)

// swagger:model CredentialsUserPayload
type CredentialsUserPayload struct {

	// Email of the user.
	//
	// required: true
	// example: dwight.schrute@dunder-mifflin.com
	Email string `json:"email"`

	// Password for the user with email.
	//
	// required: true
	// minimum: 4
	// maximum: 120
	// example: password
	Password string `json:"password"`
}

func (cup CredentialsUserPayload) Validate() error {
	return validation.ValidateStruct(&cup,
		validation.Field(&cup.Email, model.UserEmailRequiredRule...),
		validation.Field(&cup.Password, model.UserPasswordRule...),
	)
}
