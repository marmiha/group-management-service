package restapi

import (
	"group-management-api/adapter"
)

type Server struct {
	adapter.BusinessDomain
	JwtSecret string
}

func NewServer(domain *adapter.BusinessDomain, jwtSecret string) *Server {
	return &Server{
		BusinessDomain: *domain,
		JwtSecret: jwtSecret,
	}
}