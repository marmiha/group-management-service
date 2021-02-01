package modelpg

import "group-management-api/domain/model"

type UserID EntityID
type User struct {
	tableName struct{} `pg:"\"usr\",alias:u"`
	Entity

	ID           UserID `pg:"id,pk"`
	Email        string `pg:",unique"`
	Name         string
	PasswordHash string

	GroupID GroupID
	Group   Group `pg:"rel:has-one"`
}

func (u *User) MapTo(um *model.User) {
	UserToModel(u, um)
}

func (u *User) MapFrom(um *model.User) {
	ModelToUser(um, u)
}

func (u *User) ToModel() *model.User {
	um := new(model.User)
	UserToModel(u, um)
	return um
}

func NewUserFrom(um *model.User) *User {
	upg := new(User)
	upg.MapFrom(um)
	return upg
}

func NewUser(userID model.UserID) *User {
	return &User{ID: UserID(userID)}
}
