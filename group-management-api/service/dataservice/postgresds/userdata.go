package postgresds

import (
	"errors"
	"github.com/go-pg/pg/v10"
	"group-management-api/domain/model"
	"group-management-api/service/dataservice"
	"group-management-api/service/dataservice/postgresds/pgmodel"
)

// Implementation compile time check.
var _ dataservice.UserDataInterface = UserData{}

type UserData struct {
	*pg.DB
}

func (ud UserData) Create(user *model.User) error {
	userPg := pgmodel.NewUserFrom(user)

	_, err := ud.Model(userPg).Insert()
	if err != nil {
		return err
	}

	userPg.MapTo(user)
	return nil
}

func (ud UserData) Modify(user *model.User) error {
	panic("implement me")
}

func (ud UserData) Delete(id model.UserID) error {
	panic("implement me")
}

func (ud UserData) GetById(id model.UserID) (*model.User, error) {
	user := pgmodel.NewUser(id)

	err := ud.Model(user).
		Relation("Group").
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

	return user.ToModel(), nil
}

func (ud UserData) GetByEmail(email string) (*model.User, error) {
	user := new(pgmodel.User)

	err := ud.Model(user).
		Relation("Group").
		Where("email = ?", email).
		Select()

	if err != nil {
		if errors.Is(err, pg.ErrNoRows) {
			return nil, dataservice.ErrNotFound
		}
		// TODO: Error logging?
		return nil, err
	}

	return user.ToModel(), nil
}

func (ud UserData) GetList(page int, limit int) (*[]model.Group, error) {
	panic("implement me")
}

func (ud UserData) GetListAll() (*[]model.User, error) {
	panic("implement me")
}


