package dataservice

import "errors"

var (
	ErrNotFound      = errors.New("NotFound")
	ErrUserNotFound  = errors.New("UserNotFound")
	ErrGroupNotFound = errors.New("GroupNotFound")
)
