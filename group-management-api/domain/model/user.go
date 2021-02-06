package model

import validation "github.com/go-ozzo/ozzo-validation/v4"

// swagger:model UserID
type UserID EntityID

// User model which the domain uses.
// swagger:model User
type User struct {

	Entity

	// id of user
	//
	// required: true
	// example: 42
	ID           UserID `json:"id"`

	// email of the user
	//
	// required: true
	// example: dwight.schrute@gmail.com
	Email        string `json:"email"`

	// name of the user
	//
	// required: false
	// example: Dwight Schrute
	Name         string `json:"name"`

	// hash representation of password
	//
	// required: true
	// example: 5e884898da28047151d0e56f8dc6292773603d0d6aabbdd62a11ef721d1542d8
	PasswordHash string `json:"-"`

	// the group that the user is in
	//
	// required: false
	Group        *Group  `json:"group,omitempty"`
}

func (u User) Validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.Name, UserNameRule...),
		validation.Field(&u.Email, UserEmailRequiredRule...),
	)
}
