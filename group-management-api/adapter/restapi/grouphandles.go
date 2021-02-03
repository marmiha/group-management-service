package restapi

import (
	"fmt"
	"group-management-api/domain/payload"
	"net/http"
)

func (s *Server) leaveGroup(writer http.ResponseWriter, request *http.Request) {
	user := currentUserFromCtx(request)
	err := s.ManageGroup.LeaveGroup(user.ID)

	if err != nil {
		fmt.Printf("Err: %v", err)
		badRequestResponse(writer, err)
		return
	}

	successfulDeleteResponse(writer)
}

func (s *Server) joinGroup(writer http.ResponseWriter, request *http.Request) {
	var p payload.JoinGroup

	next := validatePayload(func(writer http.ResponseWriter, request *http.Request) {
		user := currentUserFromCtx(request)

		group, err := s.ManageGroup.AssignUserToGroup(user.ID, p)

		if err != nil {
			badRequestResponse(writer, err)
			return
		}

		okResponse(writer, group)
	}, &p)

	next.ServeHTTP(writer, request)
}

func (s *Server) deleteGroup(writer http.ResponseWriter, request *http.Request) {
	group := groupFromCtx(request)
	err := s.ManageGroup.DeleteGroup(group.ID)

	if err != nil {
		badRequestResponse(writer, err)
		return
	}

	successfulDeleteResponse(writer)
}

func (s *Server) modifyGroup(writer http.ResponseWriter, request *http.Request) {
	var p payload.ModifyGroupPayload

	next := validatePayload(func(writer http.ResponseWriter, request *http.Request) {
		group := groupFromCtx(request)

		group, err := s.ManageGroup.ModifyGroup(group.ID, p)

		if err != nil {
			badRequestResponse(writer, err)
			return
		}

		okResponse(writer, group)
	}, &p)

	next.ServeHTTP(writer, request)
}

func (s *Server) getGroup(writer http.ResponseWriter, request *http.Request) {
	group := groupFromCtx(request)
	okResponse(writer, group)
}

func (s *Server) getGroups(writer http.ResponseWriter, _ *http.Request) {
	groups, err := s.ListGroup.GroupsList()

	if err != nil {
		badRequestResponse(writer, err)
		return
	}

	okResponse(writer, groups)
}

func (s *Server) getUsersOfGroup(writer http.ResponseWriter, request *http.Request) {
	group := groupFromCtx(request)

	users, err :=  s.ListGroup.UsersOfGroupList(group.ID)

	if err != nil {
		badRequestResponse(writer, err)
		return
	}

	okResponse(writer, users)
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