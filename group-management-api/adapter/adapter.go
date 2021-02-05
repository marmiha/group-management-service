// Entry point for any definition of APA (gRPC, REST, graphQL).
package adapter

import (
	"group-management-api/domain/usecase"
)

type BusinessDomain struct {
	ListUser         usecase.ListUserUseCaseInterface
	ListGroup        usecase.ListGroupUseCaseInterface
	ManageGroup      usecase.ManageGroupUseCaseInterface
	ManageUser       usecase.ManageUserUseCaseInterface
	UserRegistration usecase.UserRegistrationUseCaseInterface
}