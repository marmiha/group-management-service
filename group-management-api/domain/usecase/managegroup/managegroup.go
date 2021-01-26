package managegroup

import (
	"errors"
	"group-management-api/domain/model"
	"group-management-api/domain/payload"
	"group-management-api/domain/usecase"
	"group-management-api/service/dataservices"
)

// ManageGroupUseCaseInterface compile time implementation check.
var _ usecase.ManageGroupUseCaseInterface = ManageGroupUseCase{}

type ManageGroupUseCase struct {
	GroupData dataservices.GroupDataInterface
}

func (mg ManageGroupUseCase) CreateGroup(p payload.CreateGroupPayload) (*model.Group, error) {
	group := &model.Group{
		Name: p.Name,
	}

	// Validation check.
	err := group.Validate()
	if err != nil {
		return nil, err
	}

	// Duplicate check
	_, err = mg.GroupData.GetByName(group.Name)
	if !errors.Is(err, dataservices.ErrNotFound) {
		return nil, err
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

func (mg ManageGroupUseCase) LeaveGroup(user *model.User) error {
	// Datastore fetch.
	_, err := mg.GetGroupOfUser(user.ID)
	if err != nil {
		return err
	}

	// Instruct datastore to remove user from group.
	err = mg.GroupData.LeaveGroup(user.ID)
	if err != nil {
		return err
	}

	return nil
}

func (mg ManageGroupUseCase) AssignUserToGroup(p payload.AssignUserToGroup) (*model.Group, error) {
	// Payload validation.
	err := p.Validate()
	if err != nil {
		return nil, err
	}

	// Check if user already has a group assigned.
	group, err := mg.GroupData.GetGroupOfUser(p.UserID)
	if !errors.Is(err, dataservices.ErrNotFound) {
		return nil, err
	}

	// Instruct datastore to assign user to group.
	group, err = mg.GroupData.AssignUserToGroup(p.UserID, p.GroupID)
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
