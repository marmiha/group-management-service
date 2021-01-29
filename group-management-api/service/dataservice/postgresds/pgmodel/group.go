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

func MapGroup(gm *model.Group) *Group {
	var gpg *Group
	GroupPgToGroupModel(gpg, gm)
	return gpg
}

func (g Group) MapTo(gm *model.Group) {
	GroupPgToGroupModel(&g, gm)
}

func (g Group) MapFrom(gm *model.Group) {
	GroupModelToGroupPg(gm, &g)
}
