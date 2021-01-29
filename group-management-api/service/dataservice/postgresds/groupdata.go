package postgresds

import (
	"github.com/go-pg/pg/v10"
	"group-management-api/domain/model"
)

type GroupData struct {
 	DB *pg.DB 
}

func (g GroupData) Create(user *model.User) error {
	panic("implement me")
}

func (g GroupData) Modify(user *model.User) error {
	panic("implement me")
}

func (g GroupData) Delete(id model.UserID) error {
	panic("implement me")
}

func (g GroupData) GetById(id model.UserID) (*model.User, error) {
	panic("implement me")
}

func (g GroupData) GetByEmail(email string) (*model.User, error) {
	panic("implement me")
}

func (g GroupData) GetList(page int, limit int) (*[]model.Group, error) {
	panic("implement me")
}

func (g GroupData) GetListAll() (*[]model.User, error) {
	panic("implement me")
}
