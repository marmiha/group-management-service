package restapi

import (
	"github.com/go-chi/chi"
	"group-management-api/adapter"
)

var _ adapter.ApiInterface = RestApi{}
type RestApi struct{}

func (r RestApi) SetupRouter(bd *adapter.BusinessDomain, router *chi.Mux) {
	// Pass the business logic to our server.
	server := NewServer(bd)

	// Pass the router to our server, so it sets it up.
	server.setupEndpoints(router)
}

