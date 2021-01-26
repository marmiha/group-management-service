package listuser

import (
	"group-management-api/domain/model"
	"group-management-api/domain/usecase"
	"group-management-api/service/dataservices"
)

// ListUserUseCaseInterface compile time implementation check.
var _ usecase.ListUserUseCaseInterface = ListUserUseCase{}

type ListUserUseCase struct {
	UserData dataservices.UserDataInterface
}

func (lu ListUserUseCase) Find(id model.UserID) (*model.User, error) {
	panic("implement me")
}

func (lu ListUserUseCase) UsersList() (*[]model.User, error) {
	panic("implement me")
}

