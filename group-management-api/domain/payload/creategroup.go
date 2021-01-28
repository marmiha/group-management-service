package payload

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"group-management-api/domain/model"
)

type CreateGroupPayload struct {
	Name string `json:"name"`
}

func (cgp CreateGroupPayload) Validate() error {
	return validation.ValidateStruct(&cgp,
		validation.Field(&cgp.Name, model.GroupNameRequiredRule...),
	)
}
