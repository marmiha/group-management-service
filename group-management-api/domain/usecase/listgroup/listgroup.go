package listgroup

import (
	"group-management-api/domain/model"
	"group-management-api/domain/usecase"
)

// ListGroupUseCaseInterface compile time implementation check.
var _ usecase.ListGroupUseCaseInterface = ListGroupUseCase{}

type ListGroupUseCase struct {

}

func (lg ListGroupUseCase) Find(id model.GroupID) (user *model.User, err error) {
	panic("implement me")
}

func (lg ListGroupUseCase) GroupsList() (groupList []model.Group, err error) {
	panic("implement me")
}

func (lg ListGroupUseCase) ListUsersOfGroup(id model.GroupID) (userList []model.User, err error) {
	panic("implement me")
}

