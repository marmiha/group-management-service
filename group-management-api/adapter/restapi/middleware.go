package restapi

import (
	"context"
	"errors"
	"github.com/go-chi/chi"
	"group-management-api/domain"
	"group-management-api/domain/model"
	"net/http"
	"strconv"
)

// Injects User from the JWT token into the context.
func (s *Server) WithUserAuthenticationCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// This is where we check for our token and save it inside our context.
		var tokenClaims TokenClaims
		_, err := ParseToken(r, &tokenClaims)

		// If signature is invalid or the token does not exist.
		if err != nil {
			unauthorizedResponse(w, err)
			return
		}

		// If the token is not valid (expired...).
		if err := tokenClaims.Valid(); err != nil {
			unauthorizedResponse(w, err)
			return
		}

		// Get the user from the database and insert it into the context.
		user, err := s.ListUser.Find(tokenClaims.UserID)
		if err != nil {
			if err == domain.ErrNoResult {
				unauthorizedResponse(w, ErrInvalidBearerToken)
				return
			}
			internalServerErrorResponse(w, err)
			return
		}

		ctx := context.WithValue(r.Context(), contextCurrentUserKey, user)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// Injects the Group into the context.
func (s *Server) GroupCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		group := new(model.Group)

		if stringId := chi.URLParam(r, groupIdParam); stringId != "" {
			groupID, err := strconv.ParseInt(stringId, 0, 0)
			if err != nil {
				badRequestResponse(w, err)
				return
			}

			group, err = s.ListGroup.Find(model.GroupID(groupID))
			if err != nil {
				jsonResponse(w, map[string]string{}, http.StatusBadRequest)
				return
			}
		}

		ctx := context.WithValue(r.Context(), contextGroupKey, group)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (s *Server) CurrentUserGroupCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		currentUser := currentUserFromCtx(r)

		if currentUser == nil{
			internalServerErrorResponse(w, errors.New("CurrentUserGroupCtx has to be after UserAuthCtx"))
			return
		}

		group, _ := s.ManageGroup.GetGroupOfUser(currentUser.ID)

		ctx := context.WithValue(r.Context(), contextGroupKey, group)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}


// Injects the User into the context.
func (s *Server) UserCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := new(model.User)

		if stringId := chi.URLParam(r, userIdParam); stringId != "" {
			todoId, err := strconv.ParseInt(stringId, 0, 0)
			if err != nil {
				jsonResponse(w, map[string]string{}, http.StatusBadRequest)
				return
			}

			user, err = s.ListUser.Find(model.UserID(todoId))
			if err != nil {
				notFoundResponse(w, domain.ErrNoResult)
				return
			}
		}

		ctx := context.WithValue(r.Context(), contextUserKey, user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}