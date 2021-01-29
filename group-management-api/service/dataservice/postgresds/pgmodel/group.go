package pgmodel

type GroupID EntityID
type Group struct {
	tableName struct{} `pg:"\"group\",alias:g"`
	Entity

	ID      GroupID `pg:"id,pk"`
	Name    string  `pg:",unique"`
	Members []*User `pg:"rel:has-many"`
}
