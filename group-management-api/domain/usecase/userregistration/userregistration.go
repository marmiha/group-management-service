package userregistration

import (
	"group-management-api/domain/model"
	"group-management-api/domain/payload"
	"group-management-api/domain/usecase"
)

// UserRegistrationUseCaseInterface compile time implementation check.
var _ usecase.UserRegistrationUseCaseInterface = UserRegistrationUseCase{}

type UserRegistrationUseCase struct {

}

func (ur UserRegistrationUseCase) RegisterUser(p payload.RegisterUserPayload) (user *model.User, err error) {
	panic("implement me")
}

func (ur UserRegistrationUseCase) UnregisterUser(p payload.UnregisterUserPayload) error {
	panic("implement me")
}

func (ur UserRegistrationUseCase) ChangePassword(p payload.ChangePasswordPayload) error {
	panic("implement me")
}

func (ur UserRegistrationUseCase) ValidateUserCredentials(p payload.CredentialsUserPayload) error {
	panic("implement me")
}

