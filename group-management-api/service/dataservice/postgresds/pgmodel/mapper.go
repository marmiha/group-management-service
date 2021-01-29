package pgmodel

import "group-management-api/domain/model"

func UserPgToUserModel(upg *User, um *model.User) {
	// Entity struct fields
	um.UpdatedAt = upg.UpdatedAt
	um.CreatedAt = upg.CreatedAt

	// User struct fields.
	um.ID = model.UserID(upg.ID)
	um.Email = upg.Email
	um.Name = upg.Name
	um.PasswordHash = upg.PasswordHash

	// Group conversion.
	var gm *model.Group
	GroupPgToGroupModel(&upg.Group, gm)
	um.Group = *gm
}

func GroupPgToGroupModel(gpg *Group, gm *model.Group) {
	// Entity struct fields
	gm.UpdatedAt = gpg.UpdatedAt
	gm.CreatedAt = gpg.CreatedAt

	// Group struct fields.
	gm.ID = model.GroupID(gpg.ID)
	gm.Name = gpg.Name

	// Users conversion.
	for _, upg := range gpg.Members {
		var um *model.User
		UserPgToUserModel(upg, um)
		gm.Members = append(gm.Members, um)
	}
}

func UserModelToUserPg(um *model.User, upg *User) {
	// Entity struct fields
	upg.UpdatedAt = um.UpdatedAt
	upg.CreatedAt = um.CreatedAt

	// User struct fields.
	upg.ID = UserID(um.ID)
	upg.Email = um.Email
	upg.Name = um.Name
	upg.PasswordHash = um.PasswordHash

	// Group conversion.
	var gm *Group
	GroupModelToGroupPg(&um.Group, gm)
	upg.Group = *gm
}

func GroupModelToGroupPg(gm *model.Group, gpg *Group) {
	// Entity struct fields
	gpg.UpdatedAt = gm.UpdatedAt
	gpg.CreatedAt = gm.CreatedAt

	// Group struct fields.
	gpg.ID = GroupID(gm.ID)
	gpg.Name = gm.Name

	// Users conversion.
	for _, um := range gm.Members {
		var upg *User
		UserModelToUserPg(um, upg)
		gpg.Members = append(gpg.Members, upg)
	}
}

