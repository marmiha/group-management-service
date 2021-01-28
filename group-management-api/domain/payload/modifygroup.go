package payload

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"group-management-api/domain/model"
)

type ModifyGroupPayload struct {
	Name string `json:"name"`
}

func (mgp ModifyGroupPayload) Validate() error {
	return validation.ValidateStruct(&mgp,
		validation.Field(&mgp.Name, model.GroupNameRequiredRule...),
	)
}
