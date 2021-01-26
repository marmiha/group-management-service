package model

type UserID EntityID
type User struct {
	Entity

	ID       UserID `json:"id"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Group    Group  `json:"group"`
}

func (u User) Validate() error {
	panic("implement me")
}
