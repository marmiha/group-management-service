// Entry point for any definition of APA (gRPC, REST, graphQL).
package adapter

import (
	"github.com/go-chi/chi"
	"group-management-api/domain/usecase"
)

type ApiInterface interface {
	SetupRouter(bd *BusinessDomain, router *chi.Mux)
}

type BusinessDomain struct {
	ListUser         usecase.ListUserUseCaseInterface
	ListGroup        usecase.ListGroupUseCaseInterface
	ManageGroup      usecase.ManageGroupUseCaseInterface
	ManageUser       usecase.ManageUserUseCaseInterface
	UserRegistration usecase.UserRegistrationUseCaseInterface
}