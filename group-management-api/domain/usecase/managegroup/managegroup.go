package managegroup

import (
	"errors"
	"group-management-api/domain"
	"group-management-api/domain/model"
	"group-management-api/domain/payload"
	"group-management-api/domain/usecase"
	"group-management-api/service/dataservice"
)

// ManageGroupUseCaseInterface compile time implementation check.
var _ usecase.ManageGroupUseCaseInterface = ManageGroupUseCase{}

type ManageGroupUseCase struct {
	GroupData dataservice.GroupDataInterface
}

func (mg ManageGroupUseCase) DeleteGroup(id model.GroupID) error {
	err := mg.GroupData.Delete(id)

	if err != nil {
		return err
	}

	return nil
}

func (mg ManageGroupUseCase) CreateGroup(p payload.CreateGroupPayload) (*model.Group, error) {
	group := &model.Group{
		Name: p.Name,
		Members: []*model.User{},
	}

	// Validation check.
	err := group.Validate()
	if err != nil {
		return nil, err
	}

	// Duplicate check
	groupDuplicate, err := mg.GroupData.GetByName(group.Name)
	if groupDuplicate != nil {
		return nil, domain.ErrGroupWithNameAlreadyTaken
	}

	// Group creation.
	err = mg.GroupData.Create(group)
	if err != nil {
		return nil, err
	}

	return group, nil
}

func (mg ManageGroupUseCase) ModifyGroup(id model.GroupID, p payload.ModifyGroupPayload) (*model.Group, error) {
	// Datastore fetch.
	group, err := mg.GroupData.GetById(id)
	if err != nil {
		return nil, err
	}

	// Replace name.
	group.Name = p.Name
	err = group.Validate()
	if err != nil {
		return nil, err
	}

	// Modify in datastore.
	err = mg.GroupData.Modify(group)
	if err != nil {
		return nil, err
	}

	return group, nil
}

func (mg ManageGroupUseCase) LeaveGroup(userID model.UserID) error {
	// Datastore fetch.
	group, _ := mg.GetGroupOfUser(userID)
	if group == nil {
		return nil
	}

	// Instruct datastore to remove user from group.
	err := mg.GroupData.LeaveGroup(userID)
	if err != nil {
		return err
	}

	return nil
}

func (mg ManageGroupUseCase) AssignUserToGroup(userID model.UserID, p payload.JoinGroup) (*model.Group, error) {
	// Payload validation.
	err := p.Validate()
	if err != nil {
		return nil, err
	}

	// Check if user already has a group assigned.
	group, err := mg.GroupData.GetGroupOfUser(userID)
	if err != nil && !errors.Is(err, dataservice.ErrNotFound) {
		return nil, err
	}

	if group != nil {
		return nil , domain.ErrUserAlreadyInGroup
	}


	// Instruct datastore to assign user to group.
	group, err = mg.GroupData.AssignUserToGroup(userID, p.GroupID)
	if err != nil {
		return nil, err
	}
	return group, err
}

func (mg ManageGroupUseCase) GetGroupOfUser(id model.UserID) (*model.Group, error) {
	group, err := mg.GroupData.GetGroupOfUser(id)
	if err != nil {
		return nil, err
	}
	return group, nil
}
