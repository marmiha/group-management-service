package userregistration

import (
	"golang.org/x/crypto/bcrypt"
	"group-management-api/domain"
	"group-management-api/domain/model"
	"group-management-api/domain/payload"
	"group-management-api/domain/usecase"
	"group-management-api/service/dataservice"
)

// UserRegistrationUseCaseInterface compile time implementation check.
var _ usecase.UserRegistrationUseCaseInterface = UserRegistrationUseCase{}

type UserRegistrationUseCase struct {
	UserData dataservice.UserDataInterface
}

func (ur UserRegistrationUseCase) RegisterUser(p payload.RegisterUserPayload) (*model.User, error) {
	// Unique email checkpoint.
	userExist, _ := ur.UserData.GetByEmail(p.Email)
	if userExist != nil {
		// Email is already taken.
		return nil, domain.ErrUserWithEmailAlreadyExists
	}

	// Getting the password hash.
	passwordHash, err := hashPassword(p.Password)
	if err != nil {
		return nil, err
	}

	// User struct with the payload data.
	user := &model.User{
		Name: p.Name,
		Email: p.Email,
		PasswordHash: *passwordHash,
	}

	// Try to create our user in the database.
	err = ur.UserData.Create(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (ur UserRegistrationUseCase) UnregisterUser(userID model.UserID, p payload.UnregisterUserPayload) error {
	user, err := ur.UserData.GetById(userID)
	if err != nil {
		return err
	}

	// Check if password matches the hash.
	err = checkPassword(user, p.Password)

	if err != nil {
		return domain.ErrPasswordsDoNotMatch
	}

	err = ur.UserData.Delete(userID)
	return err
}

func (ur UserRegistrationUseCase) ChangePassword(userID model.UserID, p payload.ChangePasswordPayload) (*model.User, error) {
	user, err := ur.UserData.GetById(userID)
	if err != nil {
		return nil, err
	}

	// Check if password matches the hash.
	err = checkPassword(user, p.CurrentPassword)

	if err != nil {
		return nil, domain.ErrPasswordsDoNotMatch
	}

	// Getting the new password hash.
	passwordHash, err := hashPassword(p.NewPassword)
	if err != nil {
		return nil, err
	}

	// Swap the password hash.
	user.PasswordHash = *passwordHash
	err = ur.UserData.Modify(user)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (ur UserRegistrationUseCase) ValidateUserCredentials(p payload.CredentialsUserPayload) (*model.User, error) {

	user, err := ur.UserData.GetByEmail(p.Email)
	if err != nil {
		return nil, domain.ErrInvalidLoginCredentials
	}

	// Check if password matches the hash.
	err = checkPassword(user, p.Password)

	if err != nil {
		return nil, domain.ErrInvalidLoginCredentials
	}

	return user, nil
}

func checkPassword(user *model.User, password string) error {
	bytePassword, byteHashedPassword := []byte(password), []byte(user.PasswordHash)
	return bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)
}


func hashPassword(password string) (*string, error) {
	bytePassword := []byte(password)
	passwordHash, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

	hashed := string(passwordHash)
	return &hashed, nil
}

