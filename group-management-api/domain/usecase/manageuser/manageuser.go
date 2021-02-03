package manageuser

import (
	"group-management-api/domain"
	"group-management-api/domain/model"
	"group-management-api/domain/payload"
	"group-management-api/domain/usecase"
	"group-management-api/service/dataservice"
)

// ManageUserUseCaseInterface compile time implementation check.
var _ usecase.ManageUserUseCaseInterface = ManageUserUseCase{}

type ManageUserUseCase struct {
	UserData dataservice.UserDataInterface
}

func (mu ManageUserUseCase) ModifyUserDetails(id model.UserID, p payload.ModifyUserPayload) (*model.User, error) {
	// Payload Validation
	err := p.Validate()
	if err != nil {
		return nil, err
	}

	// Find the user.
	user, err := mu.UserData.GetById(id)
	if err != nil {
		return nil, err
	}

	// Check if email modification is present.
	if p.Email != "" {
		// Check if email is already taken.
		userEmail, _ := mu.UserData.GetByEmail(p.Email)
		if userEmail != nil {
			return nil, domain.ErrUserWithEmailAlreadyExists
		}

		// Change email.
		user.Email = p.Email
	}

	// Check if name present and update if it is.
	if p.Name != "" {
		user.Name = p.Name
	}

	// Instruct the datastore to modify the user.
	err = mu.UserData.Modify(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

