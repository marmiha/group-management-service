package payload

import "group-management-api/domain/model"

type AssignUserToGroup struct {
	UserID model.UserID `json:"user_id"`
	GroupID model.GroupID `json:"group_id"`
}

func (a AssignUserToGroup) Validate() error {
	panic("implement me")
}

