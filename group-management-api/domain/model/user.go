package model

import validation "github.com/go-ozzo/ozzo-validation/v4"

type UserID EntityID
type User struct {
	Entity

	ID           UserID `json:"id"`
	Email        string `json:"email"`
	Name         string `json:"name"`
	PasswordHash string `json:"-"`
	Group        Group  `json:"group,omitempty"`
}

func (u User) Validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.Name, UserNameRule...),
		validation.Field(&u.Email, UserEmailRequiredRule...),
	)
}
