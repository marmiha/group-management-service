package postgresds

import (
	"errors"
	"github.com/go-pg/pg/v10"
	"group-management-api/domain/model"
	"group-management-api/service/dataservice"
	"group-management-api/service/dataservice/postgresds/modelpg"
)

// Implementation compile time check.
var _ dataservice.UserDataInterface = UserData{}

type UserData struct {
	*pg.DB
}

func (ud UserData) Create(user *model.User) error {
	userPg := modelpg.NewUserFrom(user)

	_, err := ud.Model(userPg).Insert()
	if err != nil {
		return err
	}

	userPg.MapTo(user)
	return nil
}

func (ud UserData) Modify(user *model.User) error {
	userPg := modelpg.NewUserFrom(user)

	_, err := ud.Model(userPg).
		WherePK().
		UpdateNotZero()

	if err != nil {
		return err
	}

	userPg.MapTo(user)
	return nil
}

func (ud UserData) Delete(id model.UserID) error {
	userPg := modelpg.NewUser(id)

	_, err := ud.Model(userPg).
		WherePK().
		Delete()

	if err != nil {
		return err
	}

	return nil
}

func (ud UserData) GetById(id model.UserID) (*model.User, error) {
	userPg := modelpg.NewUser(id)

	err := ud.Model(userPg).
		Relation("Group").
		WherePK().
		Select()

	if err != nil {
		if errors.Is(err, pg.ErrNoRows) {
			return nil, dataservice.ErrNotFound
		}
		return nil, err
	}

	return userPg.ToModel(), nil
}

func (ud UserData) GetByEmail(email string) (*model.User, error) {
	userPg := new(modelpg.User)

	err := ud.Model(userPg).
		Relation("Group").
		Where("email = ?", email).
		Select()

	if err != nil {
		if errors.Is(err, pg.ErrNoRows) {
			return nil, dataservice.ErrNotFound
		}
		return nil, err
	}

	return userPg.ToModel(), nil
}

func (ud UserData) GetListAll() (*[]model.User, error) {
	var usersPg *[]modelpg.User

	err := ud.Model(usersPg).Select()
	if err != nil {
		return nil, err
	}

	return modelpg.UsersToModels(usersPg), nil
}


