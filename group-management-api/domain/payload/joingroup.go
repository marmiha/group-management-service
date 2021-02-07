package payload

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"group-management-api/domain/model"
)

// swagger:model JoinGroup
type JoinGroup struct {

	// ID of the group to join
	//
	// required: true
	// minimum: 1
	// example: 1
	GroupID model.GroupID `json:"group_id"`
}

func (a JoinGroup) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.GroupID, validation.Required),
	)
}
