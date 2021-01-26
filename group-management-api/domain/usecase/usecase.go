package usecase

import (
	"group-management-api/domain/model"
	"group-management-api/domain/payload"
)

// Registration business logic.
type UserRegistrationUseCaseInterface interface {
	// Registration related functions.
	RegisterUser(p payload.RegisterUserPayload) (*model.User, error)
	UnregisterUser(p payload.UnregisterUserPayload) error
	// Password and access related functions.
	ChangePassword(p payload.ChangePasswordPayload) error
	ValidateUserCredentials(p payload.CredentialsUserPayload) error
}

// User management business logic.
type ManageUserUseCaseInterface interface {
	ModifyUserDetails(user *model.User) error
}

// Group management business logic.
type ManageGroupUseCaseInterface interface {
	// Basic functionalities.
	CreateGroup(p payload.CreateGroupPayload) (*model.Group, error)
	ModifyGroup(id model.GroupID, p payload.ModifyGroupPayload) (*model.Group, error)

	// User management inside groups.
	LeaveGroup(user *model.User) error
	AssignUserToGroup(group payload.AssignUserToGroup) (*model.Group, error)
	GetGroupOfUser(id model.UserID) (*model.Group, error)
}

// User listing business logic.
type ListUserUseCaseInterface interface {
	Find(id model.UserID) (*model.User, error)
	UsersList() (*[]model.User, error)
}

// Group listing business logic.
type ListGroupUseCaseInterface interface {
	Find(id model.GroupID) (*model.Group, error)
	GroupsList() (*[]model.Group, error)

	ListUsersOfGroup(id model.GroupID) (*[]model.User, error)
}
