package payload

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"group-management-api/domain/model"
)

// swagger:model ModifyUser Payload
type ModifyUserPayload struct {

	// Used for changing the users name.
	//
	// minimum: 3
	// maximum: 40
	// example: Michael Scott
	Name string `json:"name"`

	// Used for changing the users email.
	//
	// required: true
	// example: michael.scott@dunder-mifflin.com
	Email string `json:"email"`
}

func (mup ModifyUserPayload) Validate() error {
	return validation.ValidateStruct(&mup,
		// If Email was not prompted to be changed then Name should be present.
		validation.Field(
			&mup.Name,
			validation.When(validation.IsEmpty(&mup.Email), model.UserNameRequiredRule...).Else(model.UserNameRule...),
		),
		validation.Field(&mup.Email, model.UserEmailRule...),
	)
}
