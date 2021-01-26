package model

type GroupID EntityID
type Group struct {
	Entity

	ID    GroupID `json:"id"`
	Name  string  `json:"name"`
	Users []User  `json:"users"`
}

func (g Group) Validate() error {
	panic("implement me")
}
