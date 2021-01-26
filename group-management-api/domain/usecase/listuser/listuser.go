package listuser

import (
	"group-management-api/domain/model"
	"group-management-api/domain/usecase"
)

// ListUserUseCaseInterface compile time implementation check.
var _ usecase.ListUserUseCaseInterface = ListUserUseCase{}

type ListUserUseCase struct {

}

func (lu ListUserUseCase) Find(id model.UserID) (user *model.User, err error) {
	panic("implement me")
}

func (lu ListUserUseCase) UsersList() (usersList []model.User, err error) {
	panic("implement me")
}

