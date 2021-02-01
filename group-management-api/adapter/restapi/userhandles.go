package restapi

import "net/http"

func (s *Server) getUsers(writer http.ResponseWriter, _ *http.Request) {
	users, err := s.ListUser.UsersList()

	if err != nil {
		badRequestResponse(writer, err)
		return
	}

	okResponse(writer, users)
}