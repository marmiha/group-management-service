package restapi

import (
	"group-management-api/domain/payload"
	"net/http"
)

func (s *Server) getGroups(writer http.ResponseWriter, _ *http.Request) {
	groups, err := s.ListGroup.GroupsList()

	if err != nil {
		badRequestResponse(writer, err)
		return
	}

	okResponse(writer, groups)
}


func (s *Server) createGroup(writer http.ResponseWriter, request *http.Request) {
	var p payload.CreateGroupPayload

	next := validatePayload(func(writer http.ResponseWriter, request *http.Request) {
		group, err := s.ManageGroup.CreateGroup(p)

		if err != nil {
			badRequestResponse(writer, err)
			return
		}

		createdResponse(writer, group)
	}, &p)

	next.ServeHTTP(writer, request)
}