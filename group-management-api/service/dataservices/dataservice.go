package dataservices

import "group-management-api/domain/model"

type UserDataInterface interface {
	Create(user *model.User) error
	Modify(user *model.User) error
	Delete(id int) error

	GetById(id int) (*model.User, error)
	GetByEmail(email string) (*model.User, error)

	GetList(page int, limit int) (*[]model.Group, error)
}

type GroupDataInterface interface {
	Create(group *model.Group) error
	Modify(group *model.Group) error
	Delete(id int) error

	GetById(id int) (*model.Group, error)
	GetByName(name string) (*model.Group, error)

	GetList(page int, limit int) (*[]model.Group, error)
}