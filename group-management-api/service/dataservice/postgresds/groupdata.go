package postgresds

import (
	"errors"
	"github.com/go-pg/pg/v10"
	"group-management-api/domain/model"
	"group-management-api/service/dataservice"
	"group-management-api/service/dataservice/postgresds/pgmodel"
)

// Implementation compile time check.
var _ dataservice.GroupDataInterface = GroupData{}

type GroupData struct {
 	*pg.DB
}

func (gd GroupData) Create(group *model.Group) error {
	groupPg := pgmodel.NewGroupFrom(group)

	_, err := gd.Model(groupPg).Insert()
	if err != nil {
		return err
	}

	groupPg.MapTo(group)
	return nil
}

func (gd GroupData) Modify(group *model.Group) error {
	panic("implement me")
}

func (gd GroupData) Delete(id model.GroupID) error {
	panic("implement me")
}

func (gd GroupData) GetById(id model.GroupID) (*model.Group, error) {
	group := pgmodel.NewGroup(id)

	err := gd.Model(group).
		WherePK().
		Select()

	if err != nil {
		// Pass our custom error which domain handles differently.
		if errors.Is(err, pg.ErrNoRows) {
			return nil, dataservice.ErrNotFound
		}
		// TODO: Error logging?
		return nil, err
	}

	return group.ToModel(), nil
}

func (gd GroupData) GetByName(name string) (*model.Group, error) {
	group := new(pgmodel.Group)

	err := gd.Model(group).
		Where("name = ?", name).
		Select()

	if err != nil {
		if errors.Is(err, pg.ErrNoRows) {
			return nil, dataservice.ErrNotFound
		}
		// TODO: Error logging?
		return nil, err
	}

	return group.ToModel(), nil
}

func (gd GroupData) GetByUser(id model.UserID) (*model.Group, error) {
	panic("implement me")
}

func (gd GroupData) GetUsersOfGroup(id model.GroupID) (*[]model.User, error) {
	panic("implement me")
}

func (gd GroupData) GetGroupOfUser(id model.UserID) (*model.Group, error) {
	panic("implement me")
}

func (gd GroupData) GetList(page int, limit int) (*[]model.Group, error) {
	panic("implement me")
}

func (gd GroupData) GetListAll() (*[]model.Group, error) {
	panic("implement me")
}

func (gd GroupData) AssignUserToGroup(userID model.UserID, groupID model.GroupID) (*model.Group, error) {
	panic("implement me")
}

func (gd GroupData) LeaveGroup(userID model.UserID) error {
	panic("implement me")
}

