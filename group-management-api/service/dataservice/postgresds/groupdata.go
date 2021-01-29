package postgresds

import (
	"github.com/go-pg/pg/v10"
	"group-management-api/domain/model"
	"group-management-api/service/dataservice"
)

// Implementation compile time check.
var _ dataservice.GroupDataInterface = GroupData{}

type GroupData struct {
 	DB *pg.DB
}

func (g GroupData) Create(group *model.Group) error {
	panic("implement me")
}

func (g GroupData) Modify(group *model.Group) error {
	panic("implement me")
}

func (g GroupData) Delete(id model.GroupID) error {
	panic("implement me")
}

func (g GroupData) GetById(id model.GroupID) (*model.Group, error) {
	panic("implement me")
}

func (g GroupData) GetByName(name string) (*model.Group, error) {
	panic("implement me")
}

func (g GroupData) GetByUser(id model.UserID) (*model.Group, error) {
	panic("implement me")
}

func (g GroupData) GetUsersOfGroup(id model.GroupID) (*[]model.User, error) {
	panic("implement me")
}

func (g GroupData) GetGroupOfUser(id model.UserID) (*model.Group, error) {
	panic("implement me")
}

func (g GroupData) GetList(page int, limit int) (*[]model.Group, error) {
	panic("implement me")
}

func (g GroupData) GetListAll() (*[]model.Group, error) {
	panic("implement me")
}

func (g GroupData) AssignUserToGroup(user model.UserID, groupID model.GroupID) (*model.Group, error) {
	panic("implement me")
}

func (g GroupData) LeaveGroup(userID model.UserID) error {
	panic("implement me")
}

