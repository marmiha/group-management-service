package restapi

import (
	"group-management-api/adapter"
)

type Server struct {
	adapter.BusinessDomain
}

func NewServer(domain *adapter.BusinessDomain) *Server {
	return &Server{*domain}
}