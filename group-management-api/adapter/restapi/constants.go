package restapi

import "errors"

var (
	contextCurrentUserKey = "current_user"	// Request context key for the current logged in user.
	contextUserKey = "user"	// Request context key for the current logged in user.
	contextGroupKey = "group"
)

var (
	groupIdParam = "group_id"
	userIdParam = "user_id"
)

var (
	ErrInvalidBearerToken = errors.New("InvalidBearerToken")
)