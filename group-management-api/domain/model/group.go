package model

import validation "github.com/go-ozzo/ozzo-validation/v4"

// swagger:model GroupID
type GroupID EntityID

// Group model
// swagger:model Group
type Group struct {

	// swagger:allOf
	Entity

	// id of the group
	//
	// required: true
	// example: 31
	ID      GroupID `json:"id"`

	// group name
	//
	// required: true
	// example: admins
	Name    string  `json:"name"`

	// members of the group
	//
	// required: false
	// example: []
	Members []*User  `json:"users,omitempty"`
}

func (g Group) Validate() error {
	return validation.ValidateStruct(&g,
		validation.Field(&g.Name, GroupNameRequiredRule...),
	)
}
