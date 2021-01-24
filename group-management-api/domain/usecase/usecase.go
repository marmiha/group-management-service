package usecase

import (
	"group-management-api/domain/model"
	"group-management-api/domain/payload"
)
// These interfaces will be implemented for business logic.

// Registration business logic.
type RegistrationUseCaseInterface interface {
	RegisterUser(p payload.RegisterUserPayload) (user *model.User, err error)
	UnregisterUser(p payload.UnregisterUserPayload) error
	ChangePassword(p payload.ChangePasswordPayload) error
	ValidateUserCredentials(p payload.CredentialsUserPayload) error
}

// User management business logic.
type ManageUserUseCaseInterface interface {
	ModifyUserDetails(user *model.User) error
	RemoveUserFromGroup(user *model.User) error
	AddUserToGroup(user *model.User, groupID int) error
}

// Group manage business logic.
type ManageGroupUseCaseInterface interface {
	ModifyGroup(group *model.Group) error
}

// User listing business logic.
type ListUserUseCaseInterface interface {
	UsersList() (usersList []model.User, err error)
	Find(id int) (user *model.User, err error)
}

// Group listing business logic.
type ListGroupUseCaseInterface interface {
	GroupsList() (groupList []model.Group, err error)
	Find(id int) (user *model.User, err error)
}
