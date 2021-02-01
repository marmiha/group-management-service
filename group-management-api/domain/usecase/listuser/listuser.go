package listuser

import (
	"errors"
	"group-management-api/domain"
	"group-management-api/domain/model"
	"group-management-api/domain/usecase"
	"group-management-api/service/dataservice"
)

// ListUserUseCaseInterface compile time implementation check.
var _ usecase.ListUserUseCaseInterface = ListUserUseCase{}

type ListUserUseCase struct {
	UserData dataservice.UserDataInterface
}

func (lu ListUserUseCase) Find(id model.UserID) (*model.User, error) {
	user, err := lu.UserData.GetById(id)
	if err != nil {
		if errors.Is(err, dataservice.ErrNotFound) {
			return nil, domain.ErrNoResult
		}
		return nil ,err
	}
	return user, nil
}

func (lu ListUserUseCase) UsersList() ([]*model.User, error) {
	users, err := lu.UserData.GetListAll()
	if err != nil {
		return nil, err
	}
	return users, nil
}

