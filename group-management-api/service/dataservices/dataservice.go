// Interfaces that represent data sources for our domain/business models. Any implementations that will suffice this
// contract will be able to communicate with the domain logic as long as the function signatures are the same.
package dataservices

import "group-management-api/domain/model"

type UserDataInterface interface {
	Create(user *model.User) error
	Modify(user *model.User) error
	Delete(id model.UserID) error

	GetById(id model.UserID) (*model.User, error)
	GetByEmail(email string) (*model.User, error)

	GetList(page int, limit int) (*[]model.Group, error)
	GetListAll() (*[]model.Group, error)
}

type GroupDataInterface interface {
	Create(group *model.Group) error
	Modify(group *model.Group) error
	Delete(id model.GroupID) error

	GetById(id model.GroupID) (*model.Group, error)
	GetByName(name string) (*model.Group, error)
	GetByUser(id model.UserID) (*model.Group, error)

	GetList(page int, limit int) (*[]model.Group, error)
	GetListAll() (*[]model.Group, error)

	AddUserToGroup(user UserDataInterface, groupID model.GroupID) (*model.Group, error)
	RemoveUserFromGroup(user UserDataInterface)
}