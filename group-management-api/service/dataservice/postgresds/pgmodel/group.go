package pgmodel

import "group-management-api/domain/model"

type GroupID EntityID
type Group struct {
	tableName struct{} `pg:"\"group\",alias:g"`
	Entity

	ID      GroupID `pg:"id,pk"`
	Name    string  `pg:",unique"`
	Members []*User `pg:"rel:has-many"`
}

func (g Group) MapTo(gm *model.Group) {
	GroupPgToGroupModel(&g, gm)
}

func (g Group) MapFrom(gm *model.Group) {
	GroupModelToGroupPg(gm, &g)
}

func (g Group) ToModel() *model.Group {
	gm := new(model.Group)
	GroupPgToGroupModel(&g, gm)
	return gm
}

func NewGroupFrom(gm *model.Group) *Group {
	gpg := new(Group)
	GroupPgToGroupModel(gpg, gm)
	return gpg
}

func NewGroup(groupID model.GroupID) *Group {
	return &Group{ID: GroupID(groupID)}
}

