package payload

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"group-management-api/domain/model"
)

type AssignUserToGroup struct {
	UserID  model.UserID  `json:"user_id"`
	GroupID model.GroupID `json:"group_id"`
}

func (a AssignUserToGroup) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.GroupID, validation.Required),
		validation.Field(&a.UserID, validation.Required),
	)
}
