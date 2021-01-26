// Domain specific constants, business and payload validation errors.
package domain

import (
	"errors"
)

// Business logic errors.
var (
	ErrUserWithEmailAlreadyExists = errors.New("UserWithEmailAlreadyExists")
	ErrNoResult                   = errors.New("NoResult")
	ErrInvalidLoginCredentials    = errors.New("InvalidLoginCredentials")
	ErrUserNotFound               = errors.New("UserNotFound")
	ErrUserNotInGroup             = errors.New("UserNotInGroup")
)
