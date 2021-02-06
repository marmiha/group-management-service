package restapi

import "github.com/go-chi/chi"

func (s *Server) setupEndpoints(r *chi.Mux) {
	r.Route("/api/v1", func(r chi.Router) {

		// For logging the user in.
		r.Route("/login", func(r chi.Router) {

			// POST login user, gets the jwt key.
			// swagger:route POST /login loginUser
			//
			// Log in the user with email and password.
			//
			// Returns a Bearer token, if the credentials are correct.
			//
			// Responses:
			//  200: LoginResponse
			// 	400: ErrorResponse
			r.Post("/", s.loginUser)
		})

		// Routes for Users
		r.Route("/users", func(r chi.Router) {

			// GET gets all the users.
			// swagger:route GET /users getUsers
			//
			// Gets all the users.
			//
			r.Get("/", s.getUsers)

			// POST to register
			// swagger:route POST /users registerUser
			//
			// Register an user with email, name and password.
			//
			// The email has to be unique amongst the already registered members.
			//
			// Responses:
			//  200: RegisterResponse
			//  400: ErrorResponse
			r.Post("/", s.registerUser)


			// Route for logged in current user.
			r.Route("/current", func(r chi.Router) {

				// Validates the Authentication token and injects the user into request context.
				r.Use(s.WithUserAuthenticationCtx)

				// GET for getting current user.
				// swagger:route GET /users/current getSignedInUser
				//
				// Get the currently logged in user from Bearer token.
				//
				// Basically a profile fetcher.
				//
				r.Get("/", s.getCurrentUser)

				// PATCH for updating user information.
				// swagger:route PATCH /users/current modifyCurrentUser
				//
				// Modify user details, which are email and the name.
				//
				// The email has to be unique amongst the already registered users. At least one of these parameters have to be supplied.
				//
				//
				r.Patch("/", s.modifyUser)

				// DELETE for unregistering.
				// swagger:route DELETE /users/current unregisterCurrentUser
				//
				// Unregister the user that is denoted from the Bearer token.
				//
				// Supply the api with email and current password, so that that the user confirms his choice.
				//
				r.Delete("/", s.unregisterUser)

				// Routes for current user group.
				r.Route("/group", func(r chi.Router) {

					// This will get the group from Datastore and put it into request context.
					r.Use(s.CurrentUserGroupCtx)

					// GET for fetching the current group.
					// swagger:route GET /users/current/group getCurrentUserGroup
					//
					// Get the group from the currently logged in user.
					//
					// Returns bad request if the user has not joined a group yet.
					//
					r.Get("/", s.getGroup)

					// POST for joining a group.
					// swagger:route POST /users/current/group joinGroup
					//
					// Join a group denoted by group_id.
					//
					// The group should exist and the user should not be in a group already. If these circumstances are not respected an error will be returned.
					//
					r.Post("/", s.joinGroup)

					// DELETE for leaving the current group.
					// swagger:route DELETE /users/current/group leaveGroup
					//
					// Leave the current group.
					//
					// In any instance the API returns a successful delete response.
					//
					r.Delete("/", s.leaveGroup)
				})
			})

			// Routes for specific user.
			r.Route("/{" + userIdParam + "}", func(r chi.Router) {
				// Injects the user with user_id into request context.
				r.Use(s.UserCtx)

				// GET a user with user_id.
				// swagger:route GET /users/{user_id} getUser
				//
				// Get the user which is denoted by user_id.
				//
				r.Get("/", s.getUser)
			})
		})

		// Routes for Groups.
		r.Route("/groups", func(r chi.Router) {

			// GET gets all the groups.
			// swagger:route GET /groups getGroups
			//
			// Gets all the groups.
			//
			r.Get("/", s.getGroups)

			// POST to create a group.
			// swagger:route POST /groups createGroup
			//
			// Creates a group with a name, it that is not already taken.
			//
			r.Post("/", s.createGroup)

			r.Route("/{" + groupIdParam + "}", func(r chi.Router) {

				// Injects the group with group_id into request context.
				r.Use(s.GroupCtx)

				// GET to fetch the group.
				// swagger:route GET /groups/{group_id} getGroup
				//
				// Get a group which is denoted by group_id.
				//
				r.Get("/", s.getGroup)

				// DELETE to delete the group.
				// swagger:route DELETE /groups/{group_id} deleteGroup
				//
				// Delete the group denoted by group_id.
				//
				r.Delete("/", s.deleteGroup)

				// PATCH to modify group name.
				// swagger:route PATCH /groups/{group_id} modifyGroup
				//
				// Change the group name.
				//
				// If the name is taken it will return a bad request error with explanation.
				//
				r.Patch("/", s.modifyGroup)

				r.Route("/users", func(r chi.Router) {

					// GET all members of a group.
					// swagger:route GET /groups/{group_id}/users getMembersOfGroup
					//
					// Gets all the members of a group denoted by user_id.
					//
					r.Get("/", s.getUsersOfGroup)
				})
			})
		})
	})
}