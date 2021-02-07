package payload

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"group-management-api/domain/model"
)

// swagger:model ModifyGroupPayload
type ModifyGroupPayload struct {

	// Used to change the groups name.
	//
	// minimum: 3
	// maximum: 40
	// required: true
	// example: assistants to the regional manager
	Name string `json:"name"`
}

func (mgp ModifyGroupPayload) Validate() error {
	return validation.ValidateStruct(&mgp,
		validation.Field(&mgp.Name, model.GroupNameRequiredRule...),
	)
}
