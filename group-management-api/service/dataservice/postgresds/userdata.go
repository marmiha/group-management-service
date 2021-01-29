package postgresds

import (
	"github.com/go-pg/pg/v10"
	"group-management-api/domain/model"
)

type UserData struct {
	DB *pg.DB 
}

func (u UserData) Create(user *model.User) error {
	panic("implement me")
}

func (u UserData) Modify(user *model.User) error {
	panic("implement me")
}

func (u UserData) Delete(id model.UserID) error {
	panic("implement me")
}

func (u UserData) GetById(id model.UserID) (*model.User, error) {

}

func (u UserData) GetByEmail(email string) (*model.User, error) {
	panic("implement me")
}

func (u UserData) GetList(page int, limit int) (*[]model.Group, error) {
	panic("implement me")
}

func (u UserData) GetListAll() (*[]model.User, error) {
	panic("implement me")
}


