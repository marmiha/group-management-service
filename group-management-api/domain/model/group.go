package model

type Group struct {
	Entity
	Name  string `json:"name"`
	Users []User `json:"users"`
}

func (g Group) Validate() error {
	panic("implement me")
}
