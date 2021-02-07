package model

import validation "github.com/go-ozzo/ozzo-validation/v4"

// swagger:model GroupID
// example: 1
type GroupID EntityID

// Group model
// swagger:model Group
type Group struct {

	// swagger:allOf
	Entity

	// id of the group
	//
	// required: true
	// minimum: 1
	// example: 1
	ID GroupID `json:"id"`

	// group name
	//
	// required: true
	// minimum: 3
	// minimum: 40
	// example: admins
	Name string `json:"name"`

	// members of the group
	//
	// example: []
	Members []*User `json:"users,omitempty"`
}

func (g Group) Validate() error {
	return validation.ValidateStruct(&g,
		validation.Field(&g.Name, GroupNameRequiredRule...),
	)
}
