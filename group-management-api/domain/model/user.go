package model

type User struct {
	Entity
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Group    Group  `json:"group"`
}

func (u User) Validate() error {
	panic("implement me")
}
