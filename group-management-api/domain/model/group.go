package model

import validation "github.com/go-ozzo/ozzo-validation/v4"

type GroupID EntityID
type Group struct {
	Entity

	ID    GroupID `json:"id"`
	Name  string  `json:"name"`
	Users []User  `json:"users"`
}

func (g Group) Validate() error {
	return validation.ValidateStruct(&g,
		validation.Field(&g.Name, GroupNameRequiredRule...),
	)
}
