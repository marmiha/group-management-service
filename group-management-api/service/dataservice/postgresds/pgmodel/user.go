package pgmodel

type UserID EntityID
type User struct {
	tableName struct{} `pg:"\"users\",alias:u"`
	Entity

	ID           UserID `pg:"id,pk"`
	Email        string `pg:",unique"`
	Name         string
	PasswordHash string
	Group        Group  `pg:"rel:has-one"`
}
