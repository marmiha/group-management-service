package listgroup

import (
	"group-management-api/domain/model"
	"group-management-api/domain/usecase"
	"group-management-api/service/dataservice"
)

// ListGroupUseCaseInterface compile time implementation check.
var _ usecase.ListGroupUseCaseInterface = ListGroupUseCase{}

type ListGroupUseCase struct {
	GroupData dataservice.GroupDataInterface
}

func (lg ListGroupUseCase) Find(id model.GroupID) (*model.Group, error) {
	group, err := lg.GroupData.GetById(id)
	if err != nil {
		return nil, err
	}
	return group, nil
}

func (lg ListGroupUseCase) GroupsList() ([]*model.Group, error) {
	groups, err := lg.GroupData.GetListAll()
	if err != nil {
		return nil, err
	}
	return groups, nil
}

func (lg ListGroupUseCase) UsersOfGroupList(id model.GroupID) ([]*model.User, error) {
	users, err := lg.GroupData.GetUsersOfGroup(id)
	if err != nil {
		return nil, err
	}
	return users, nil
}



