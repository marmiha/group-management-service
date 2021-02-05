package restapi

import (
	"errors"
	"github.com/go-chi/chi"
	"group-management-api/adapter"
)


func SetupRouterForRestAdapter(bd *adapter.BusinessDomain, router *chi.Mux, jwtSecret string) error{
	if jwtSecret == "" {
		return errors.New("jwt token signing secret should not be an empty string")
	}
	// Pass the business logic to our server.
	server := NewServer(bd, jwtSecret)

	// Pass the router to our server, so it sets it up.
	server.setupEndpoints(router)
	return nil
}

