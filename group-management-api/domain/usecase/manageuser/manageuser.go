package manageuser

import (
	"group-management-api/domain/model"
	"group-management-api/domain/usecase"
)

// ManageUserUseCaseInterface compile time implementation check.
var _ usecase.ManageUserUseCaseInterface = ManageUserUseCase{}

type ManageUserUseCase struct {

}

func (mu ManageUserUseCase) ModifyUserDetails(user *model.User) error {
	panic("implement me")
}

