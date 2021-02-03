package modelpg

import (
	"group-management-api/domain/model"
)

type GroupID EntityID
type Group struct {
	tableName struct{} `pg:"\"grp\",alias:g"`
	Entity

	ID      GroupID `pg:"id,pk"`
	Name    string  `pg:",unique"`

	Members []*User `pg:"rel:has-many,join_fk:group_id"`
}

func (g *Group) MapTo(gm *model.Group) {
	GroupToModel(g, gm)
}

func (g *Group) MapFrom(gm *model.Group) {
	ModelToGroup(gm, g)
}

func (g *Group) ToModel() *model.Group {
	if g == nil {
		return nil
	}

	gm := new(model.Group)
	GroupToModel(g, gm)
	return gm
}

func NewGroupFrom(gm *model.Group) *Group {
	if gm == nil {
		return nil
	}

	gpg := new(Group)
	gpg.MapFrom(gm)
	return gpg
}

func NewGroup(groupID model.GroupID) *Group {
	return &Group{ID: GroupID(groupID)}
}

