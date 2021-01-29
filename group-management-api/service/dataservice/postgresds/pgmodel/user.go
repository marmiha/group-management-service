package pgmodel

import "group-management-api/domain/model"

type UserID EntityID
type User struct {
	tableName struct{} `pg:"\"users\",alias:u"`
	Entity

	ID           UserID `pg:"id,pk"`
	Email        string `pg:",unique"`
	Name         string
	PasswordHash string
	Group        Group `pg:"rel:has-one"`
}

func MapUser(um *model.User) *User {
	var upg *User
	upg.MapFrom(um)
	return upg
}

func (u User) MapTo(um *model.User) {
	UserPgToUserModel(&u, um)
}

func (u User) MapFrom(um *model.User) {
	UserModelToUserPg(um, &u)
}
