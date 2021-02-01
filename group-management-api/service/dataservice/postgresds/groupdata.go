package postgresds

import (
	"errors"
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"group-management-api/domain"
	"group-management-api/domain/model"
	"group-management-api/service/dataservice"
	"group-management-api/service/dataservice/postgresds/modelpg"
)

// Implementation compile time check.
var _ dataservice.GroupDataInterface = GroupData{}

type GroupData struct {
	*pg.DB
}

func (gd GroupData) Create(group *model.Group) error {
	groupPg := modelpg.NewGroupFrom(group)

	_, err := gd.Model(groupPg).
		Returning("*").
		Insert()

	if err != nil {
		return err
	}

	groupPg.MapTo(group)
	return nil
}

func (gd GroupData) Modify(group *model.Group) error {
	groupPg := modelpg.NewGroupFrom(group)

	_, err := gd.Model(groupPg).
		WherePK().
		Returning("*").
		UpdateNotZero()

	if err != nil {
		return err
	}

	groupPg.MapTo(group)
	return nil
}

func (gd GroupData) Delete(id model.GroupID) error {
	groupPg := modelpg.NewGroup(id)

	_, err := gd.Model(groupPg).
		WherePK().
		Delete()

	if err != nil {
		return err
	}
	return nil
}

func (gd GroupData) GetById(id model.GroupID) (*model.Group, error) {
	groupPg := modelpg.NewGroup(id)

	err := gd.Model(groupPg).
		WherePK().
		Returning("*").
		Select()

	if err != nil {
		// Pass our custom error which domain handles differently.
		if errors.Is(err, pg.ErrNoRows) {
			return nil, dataservice.ErrNotFound
		}
		return nil, err
	}

	return groupPg.ToModel(), nil
}

func (gd GroupData) GetByName(name string) (*model.Group, error) {
	groupPg := new(modelpg.Group)

	err := gd.Model(groupPg).
		Where("name = ?", name).
		Returning("*").
		Select()

	if err != nil {
		if errors.Is(err, pg.ErrNoRows) {
			return nil, dataservice.ErrNotFound
		}
		return nil, err
	}

	return groupPg.ToModel(), nil
}

func (gd GroupData) GetByUser(id model.UserID) (*model.Group, error) {
	groupPg := new(modelpg.Group)

	err := gd.Model(groupPg).
		Relation("Members", func(query *orm.Query) (*orm.Query, error) {
			return query.Where("id = ?", id), nil
		}).
		Returning("*").
		First()

	if err != nil {
		if errors.Is(err, pg.ErrNoRows) {
			return nil, domain.ErrNoResult
		}
		return nil, err
	}

	return groupPg.ToModel(), nil
}

func (gd GroupData) GetUsersOfGroup(id model.GroupID) ([]*model.User, error) {
	groupPg := modelpg.NewGroup(id)

	err := gd.Model(groupPg).
		WherePK().
		Relation("Members").
		Returning("*").
		Select()

	if err != nil {
		if errors.Is(err, pg.ErrNoRows) {
			return nil, dataservice.ErrNotFound
		}
		return nil, err
	}

	return modelpg.UsersToModels(&groupPg.Members), nil
}

func (gd GroupData) GetGroupOfUser(id model.UserID) (*model.Group, error) {
	userPg := modelpg.NewUser(id)

	err := gd.Model(userPg).
		WherePK().
		Relation("Group").
		Returning("*").
		Select()

	if err != nil {
		if errors.Is(err, pg.ErrNoRows) {
			return nil, dataservice.ErrNotFound
		}
		return nil, err
	}

	return userPg.Group.ToModel(), nil
}

func (gd GroupData) GetListAll() ([]*model.Group, error) {
	groupsPg := &[]*modelpg.Group{}

	err := gd.Model(groupsPg).
		Returning("*").
		Select()

	if err != nil {
		return nil, err
	}

	return modelpg.GroupsToModels(groupsPg), nil
}

func (gd GroupData) AssignUserToGroup(userID model.UserID, groupID model.GroupID) (*model.Group, error) {
	userPg := modelpg.NewUser(userID)
	userPg.GroupID = modelpg.GroupID(groupID)

	_, err := gd.Model(userPg).
		WherePK().
		Relation("Group").
		Returning("*").
		UpdateNotZero(userPg)

	if err != nil {
		return nil, err
	}

	return userPg.Group.ToModel(), nil
}

func (gd GroupData) LeaveGroup(userID model.UserID) error {
	userPg := modelpg.NewUser(userID)

	_, err := gd.Model(userPg).
		WherePK().
		Set("group_id = NULL").
		Returning("*").
		Update()

	if err != nil {
		return err
	}

	return nil
}
