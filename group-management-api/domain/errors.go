// Domain specific constants, business and payload validation errors.
package domain

import (
	"errors"
)

// Business logic errors.
var (
	ErrUserWithEmailAlreadyExists = errors.New("UserWithEmailAlreadyExists")
	ErrInvalidLoginCredentials    = errors.New("InvalidLoginCredentials")
	ErrPasswordsDoNotMatch        = errors.New("PasswordsDoNotMatch")

	ErrNoResult           = errors.New("NoResult")
	ErrUserAlreadyInGroup = errors.New("UserAlreadyInGroup")

	ErrGroupWithNameAlreadyTaken = errors.New("GroupWithNameAlreadyTaken")
)
