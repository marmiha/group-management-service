package modelpg

import (
	"group-management-api/domain/model"
)

// Map User to model.User.
func UserToModel(upg *User, um *model.User) {
	if um == nil {
		return
	}
	// Entity struct fields
	um.UpdatedAt = upg.UpdatedAt
	um.CreatedAt = upg.CreatedAt

	// User struct fields.
	um.ID = model.UserID(upg.ID)
	um.Email = upg.Email
	um.Name = upg.Name
	um.PasswordHash = upg.PasswordHash

	// Group conversion.
	if upg.Group != nil {
		gm := new(model.Group)
		GroupToModel(upg.Group, gm)
		um.Group = gm
	}
}

// Map Group to model.Group.
func GroupToModel(gpg *Group, gm *model.Group) {
	if gm == nil {
		return
	}

	// Entity struct fields
	gm.UpdatedAt = gpg.UpdatedAt
	gm.CreatedAt = gpg.CreatedAt

	// Group struct fields.
	gm.ID = model.GroupID(gpg.ID)
	gm.Name = gpg.Name

	// Users conversion.
	for _, upg := range gpg.Members {
		um := new(model.User)
		UserToModel(upg, um)
		gm.Members = append(gm.Members, um)
	}
}

// Map model.User to User.
func ModelToUser(um *model.User, upg *User) {
	if upg == nil {
		return
	}
	// Entity struct fields
	upg.UpdatedAt = um.UpdatedAt
	upg.CreatedAt = um.CreatedAt

	// User struct fields.
	upg.ID = UserID(um.ID)
	upg.Email = um.Email
	upg.Name = um.Name
	upg.PasswordHash = um.PasswordHash

	// Group conversion.
	if um.Group != nil {
		gm := new(Group)
		ModelToGroup(um.Group, gm)
		upg.Group = gm
	}
}

// Map model.Group to Group.
func ModelToGroup(gm *model.Group, gpg *Group) {
	if gpg == nil {
		return
	}
	// Entity struct fields
	gpg.UpdatedAt = gm.UpdatedAt
	gpg.CreatedAt = gm.CreatedAt

	// Group struct fields.
	gpg.ID = GroupID(gm.ID)
	gpg.Name = gm.Name

	// Users conversion.
	for _, um := range gm.Members {
		upg := new(User)
		ModelToUser(um, upg)
		gpg.Members = append(gpg.Members, upg)
	}
}

// Map *[]Group to *[]model.Group
func GroupsToModels(gpgs *[]*Group) []*model.Group{
	gms := &[]*model.Group{}
	for _, gpg := range *gpgs {
		gm := new(model.Group)
		GroupToModel(gpg, gm)
		*gms = append(*gms, gm)
	}
	return *gms
}

// Map *[]User to *[]model.User
func UsersToModels(upgs *[]*User) []*model.User{
	ums := &[]*model.User{}
	for _, upg := range *upgs {
		um := new(model.User)
		UserToModel(upg, um)
		*ums = append(*ums, um)
	}
	return *ums
}