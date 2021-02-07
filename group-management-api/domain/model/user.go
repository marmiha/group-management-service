package model

import validation "github.com/go-ozzo/ozzo-validation/v4"

// swagger:model UserID
// example: 1
type UserID EntityID

// User model
// swagger:model User
type User struct {

	// swagger:allOf
	Entity

	// id of user
	//
	// minimum: 1
	// required: true
	// example: 1
	ID UserID `json:"id"`

	// email of the user
	//
	// required: true
	// example: dwight.schrute@gmail.com
	// swagger:strfmt email
	Email string `json:"email"`

	// name of the user
	//
	// minimum: 3
	// maximum: 40
	// required: false
	// example: Dwight Schrute
	Name string `json:"name"`

	// hash representation of password
	//
	// required: true
	// example: 5e884898da28047151d0e56f8dc6292773603d0d6aabbdd62a11ef721d1542d8
	PasswordHash string `json:"-"`

	// the group that the user is in
	//
	// required: false
	Group *Group `json:"group,omitempty"`
}

func (u User) Validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.Name, UserNameRule...),
		validation.Field(&u.Email, UserEmailRequiredRule...),
	)
}
