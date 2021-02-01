package payload

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"group-management-api/domain/model"
)

type JoinGroup struct {
	GroupID model.GroupID `json:"group_id"`
}

func (a JoinGroup) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.GroupID, validation.Required),
	)
}
