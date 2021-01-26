package managegroup

import (
	"group-management-api/domain/model"
	"group-management-api/domain/payload"
	"group-management-api/domain/usecase"
)

// ManageGroupUseCaseInterface compile time implementation check.
var _ usecase.ManageGroupUseCaseInterface = ManageGroupUseCase{}

type ManageGroupUseCase struct {

}

func (mg ManageGroupUseCase) CreateGroup(p payload.CreateGroupPayload) (group *model.Group, err error) {
	panic("implement me")
}

func (mg ManageGroupUseCase) ModifyGroup(id model.GroupID, p payload.ModifyGroupPayload) (group *model.Group, err error) {
	panic("implement me")
}

func (mg ManageGroupUseCase) RemoveUserFromGroup(user *model.User) error {
	panic("implement me")
}

func (mg ManageGroupUseCase) AddUserToGroup(user *model.User, groupID model.GroupID) error {
	panic("implement me")
}

