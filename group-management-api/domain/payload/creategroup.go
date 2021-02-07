package payload

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"group-management-api/domain/model"
)

// swagger:model CreateGroupPayload
type CreateGroupPayload struct {

	// Name of the group.
	//
	// required: true
	// minimum: 3
	// maximum: 40
	// example: regional managers
	Name string `json:"name"`
}

func (cgp CreateGroupPayload) Validate() error {
	return validation.ValidateStruct(&cgp,
		validation.Field(&cgp.Name, model.GroupNameRequiredRule...),
	)
}
