// Interfaces related to our business logic. These should be implemented in business logic implementations based on
// the function.
package usecase

import (
	"group-management-api/domain/model"
	"group-management-api/domain/payload"
)
// These interfaces will be implemented for business logic.

// Registration business logic.
type RegistrationUseCaseInterface interface {
	// Registration related functions.
	RegisterUser(p payload.RegisterUserPayload) (user *model.User, err error)
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
	CreateGroup(p payload.CreateGroupPayload) (group *model.Group, err error)
	ModifyGroup(id model.GroupID, p payload.ModifyGroupPayload) (group *model.Group, err error)

	// User management inside groups.
	RemoveUserFromGroup(user *model.User) error
	AddUserToGroup(user *model.User, groupID model.GroupID) error
}

// User listing business logic.
type ListUserUseCaseInterface interface {
	Find(id model.UserID) (user *model.User, err error)
	UsersList() (usersList []model.User, err error)
}

// Group listing business logic.
type ListGroupUseCaseInterface interface {
	Find(id model.GroupID) (user *model.User, err error)
	GroupsList() (groupList []model.Group, err error)

	ListUsersOfGroup(id model.GroupID) (userList []model.User, err error)
}
