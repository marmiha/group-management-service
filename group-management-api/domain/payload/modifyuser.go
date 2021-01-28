package payload

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"group-management-api/domain/model"
)

type ModifyUserPayload struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (mup ModifyUserPayload) Validate() error {
	return validation.ValidateStruct(&mup,
		// If Email was not prompted to be changed then Name should be present.
		validation.Field(
			&mup.Name,
			validation.When(validation.IsEmpty(&mup.Email), model.UserNameRequiredRule...).Else(model.UserNameRule...)
		),
		validation.Field(&mup.Email, model.UserEmailRule...),
	)
}
