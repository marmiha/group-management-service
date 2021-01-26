package listgroup

import (
	"group-management-api/domain/model"
	"group-management-api/domain/usecase"
)

// ListGroupUseCaseInterface compile time implementation check.
var _ usecase.ListGroupUseCaseInterface = ListGroupUseCase{}

type ListGroupUseCase struct {

}

func (lg ListGroupUseCase) Find(id model.GroupID) (*model.User, error) {
	panic("implement me")
}

func (lg ListGroupUseCase) GroupsList() (*[]model.Group, error) {
	panic("implement me")
}

func (lg ListGroupUseCase) ListUsersOfGroup(id model.GroupID) (*[]model.User, error) {
	panic("implement me")
}



