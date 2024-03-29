package usecase

import (
	"group-management-api/domain/model"
	"group-management-api/domain/payload"
)

// Registration business logic.
type UserRegistrationUseCaseInterface interface {
	// Registration related functions.
	RegisterUser(p payload.RegisterUserPayload) (*model.User, error)
	UnregisterUser(userID model.UserID, p payload.UnregisterUserPayload) error
	// PasswordHash and access related functions.
	ChangePassword(userID model.UserID, p payload.ChangePasswordPayload) (*model.User, error)
	ValidateUserCredentials(p payload.CredentialsUserPayload) (*model.User, error)
}

// User management business logic.
type ManageUserUseCaseInterface interface {
	ModifyUserDetails(id model.UserID, p payload.ModifyUserPayload) (*model.User, error)
}

// User listing business logic.
type ListUserUseCaseInterface interface {
	Find(id model.UserID) (*model.User, error)
	UsersList() ([]*model.User, error)
}

// Group management business logic.
type ManageGroupUseCaseInterface interface {
	// Basic functionalities.
	CreateGroup(p payload.CreateGroupPayload) (*model.Group, error)
	ModifyGroup(id model.GroupID, p payload.ModifyGroupPayload) (*model.Group, error)
	DeleteGroup(id model.GroupID) error

	// User management inside groups.
	LeaveGroup(userID model.UserID) error
	AssignUserToGroup(userID model.UserID, group payload.JoinGroup) (*model.Group, error)
	GetGroupOfUser(id model.UserID) (*model.Group, error)
}

// Group listing business logic.
type ListGroupUseCaseInterface interface {
	Find(id model.GroupID) (*model.Group, error)
	GroupsList() ([]*model.Group, error)

	UsersOfGroupList(id model.GroupID) ([]*model.User, error)
}
