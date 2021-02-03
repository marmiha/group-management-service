package restapi

import (
	"group-management-api/domain/model"
	"group-management-api/domain/payload"
	"net/http"
)

type RegisterResponse struct {
	User  model.User `json:"user"`
	Token string     `json:"token"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

func (s *Server) getCurrentUser(writer http.ResponseWriter, request *http.Request) {
	user := currentUserFromCtx(request)
	okResponse(writer, user)
}


func (s *Server) getUser(writer http.ResponseWriter, request *http.Request) {
	user := userFromCtx(request)

	if user == nil {
		jsonResponse(writer, nil, http.StatusNotFound)
		return
	}

	okResponse(writer, user)
}

func (s *Server) registerUser(writer http.ResponseWriter, request *http.Request) {
	var p payload.RegisterUserPayload

	next := validatePayload(func(writer http.ResponseWriter, request *http.Request) {
		user, err := s.UserRegistration.RegisterUser(p)

		if err != nil {
			badRequestResponse(writer, err)
			return
		}

		token, err := GenerateToken(user.ID)
		if err != nil {
			badRequestResponse(writer, err)
			return
		}

		response := RegisterResponse{
			User:  *user,
			Token: *token,
		}

		createdResponse(writer, response)
	}, &p)

	next.ServeHTTP(writer, request)
}

func (s *Server) modifyUser(writer http.ResponseWriter, request *http.Request) {
	var p payload.ModifyUserPayload

	next := validatePayload(func(writer http.ResponseWriter, request *http.Request) {
		user := currentUserFromCtx(request)
		user, err := s.ManageUser.ModifyUserDetails(user.ID, p)

		if err != nil {
			badRequestResponse(writer, err)
			return
		}

		okResponse(writer, user)
	}, &p)

	next.ServeHTTP(writer, request)
}

func (s *Server) unregisterUser(writer http.ResponseWriter, request *http.Request) {
	var p payload.UnregisterUserPayload

	next := validatePayload(func(writer http.ResponseWriter, request *http.Request) {
		user := currentUserFromCtx(request)
		err := s.UserRegistration.UnregisterUser(user.ID, p)

		if err != nil {
			badRequestResponse(writer, err)
			return
		}

		successfulDeleteResponse(writer)
	}, &p)

	next.ServeHTTP(writer, request)
}

func (s *Server) loginUser(writer http.ResponseWriter, request *http.Request) {
	var p payload.CredentialsUserPayload

	next := validatePayload(func(writer http.ResponseWriter, request *http.Request) {
		user, err := s.UserRegistration.ValidateUserCredentials(p)

		if err != nil {
			badRequestResponse(writer, err)
			return
		}

		token, err := GenerateToken(user.ID)
		if err != nil {
			badRequestResponse(writer, err)
			return
		}

		response := LoginResponse{
			Token: *token,
		}

		createdResponse(writer, response)
	}, &p)

	next.ServeHTTP(writer, request)
}

func (s *Server) getUsers(writer http.ResponseWriter, _ *http.Request) {
	users, err := s.ListUser.UsersList()

	if err != nil {
		badRequestResponse(writer, err)
		return
	}

	okResponse(writer, users)
}


