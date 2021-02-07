package restapi

// These are for go-swagger OpenAPI generation.

// swagger:parameters getUser
type UserIdParam struct {
	// A specific user denoted by the user_id.
	//
	// minimum: 1
	// required: true
	// in: path
	UserID int64 `json:"user_id"`
}


// swagger:parameters getGroup deleteGroup modifyGroup getMembersOfGroup
type GroupIdParam struct {
	// A specific group denoted by the group_id.
	//
	// minimum: 1
	// required: true
	// in: path
	GroupID int64 `json:"group_id"`
}