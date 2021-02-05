package factory

import (
	"group-management-api/app/container"
	"group-management-api/domain/usecase"
	"group-management-api/domain/usecase/listgroup"
	"group-management-api/domain/usecase/listuser"
	"group-management-api/domain/usecase/managegroup"
	"group-management-api/domain/usecase/manageuser"
	"group-management-api/domain/usecase/userregistration"
)

// Singleton pattern factory for our use case factories.
// These don't need to be thread safe as the constructions will be called on single thread.
// Currently there is only one implementation for each use case, so we just return that one but we could based on the
// config parameters return some other. Pluggable stuff and what not :).

var luuc usecase.ListUserUseCaseInterface
func GetListUserUseCase(c *container.Container) (usecase.ListUserUseCaseInterface, error){
	if luuc == nil {
		tempLuuc, err := GetUserDataService(c)
		if err != nil {
			return nil, err
		}
		luuc = listuser.ListUserUseCase{UserData: tempLuuc}
	}
	return luuc, nil
}

var lguc usecase.ListGroupUseCaseInterface
func GetListGroupUseCase(c *container.Container) (usecase.ListGroupUseCaseInterface, error){
	if lguc == nil {
		tempGds, err := GetGroupDataService(c)
		if err != nil {
			return nil, err
		}
		lguc = listgroup.ListGroupUseCase{GroupData: tempGds}
	}
	return lguc, nil
}

var mguc usecase.ManageGroupUseCaseInterface
func GetManageGroupUseCase(c *container.Container) (usecase.ManageGroupUseCaseInterface, error){
	if mguc == nil {
		tempGds, err := GetGroupDataService(c)
		if err != nil {
			return nil, err
		}
		mguc = managegroup.ManageGroupUseCase{GroupData: tempGds}
	}
	return mguc, nil
}

var muuc usecase.ManageUserUseCaseInterface
func GetManageUserUseCase(c *container.Container) (usecase.ManageUserUseCaseInterface, error){
	if muuc == nil {
		tempUds, err := GetUserDataService(c)
		if err != nil {
			return nil, err
		}
		muuc = manageuser.ManageUserUseCase{UserData: tempUds}
	}
	return muuc, nil
}

var uruc usecase.UserRegistrationUseCaseInterface
func GetUserRegistrationUseCase(c *container.Container) (usecase.UserRegistrationUseCaseInterface, error){
	if uruc == nil {
		tempUds, err := GetUserDataService(c)
		if err != nil {
			return nil, err
		}

		uruc = userregistration.UserRegistrationUseCase{UserData: tempUds}
	}
	return uruc, nil
}
