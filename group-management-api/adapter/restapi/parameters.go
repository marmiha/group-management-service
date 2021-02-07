package restapi

import "group-management-api/domain/payload"

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


// swagger:parameters changeCurrentUserPassword
type ChangePasswordParam struct {

	// in: body
	ChangePasswordPayload payload.ChangePasswordPayload
}

// swagger:parameters createGroup
type CreateGroupParam struct {

	// in: body
	CreateGroupPayload payload.CreateGroupPayload
}

// swagger:parameters loginUser
type CredentialsUserParam struct {

	// in: body
	CredentialsUserPayload payload.CredentialsUserPayload
}

// swagger:parameters joinGroup
type JoinGroupParam struct {

	// in: body
	JoinGroupPayload payload.JoinGroup
}

// swagger:parameters modifyGroup
type ModifyGroupParam struct {

	// in: body
	Payload payload.ModifyGroupPayload
}

// swagger:parameters modifyCurrentUser
type ModifyUserParam struct {

	// in: body
	ModifyUserPayload payload.ModifyUserPayload
}

// swagger:parameters registerUser
type RegisterUserParam struct {

	// in: body
	RegisterUserPayload payload.RegisterUserPayload
}

// swagger:parameters unregisterCurrentUser
type UnregisterUserParam struct {

	// in: body
	UnregisterUserPayload payload.UnregisterUserPayload
}