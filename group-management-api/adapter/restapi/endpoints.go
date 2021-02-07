package restapi

import "github.com/go-chi/chi"

func (s *Server) setupEndpoints(r *chi.Mux) {
	r.Route("/api/v1", func(r chi.Router) {

		// For logging the user in.
		r.Route("/login", func(r chi.Router) {

			// POST login user, gets the jwt key.
			// swagger:route POST /login auth loginUser
			//
			// Log in the user with email and password.
			//
			// Returns a Bearer token, if the credentials are correct.
			//
			// Responses:
			//   200: LoginResponse User token used for authentication.
			// 	 400: ErrorResponse Invalid credentials error.
			//
			r.Post("/", s.loginUser)
		})

		// Routes for Users
		r.Route("/users", func(r chi.Router) {

			// GET gets all the users.
			// swagger:route GET /users users getUsers
			//
			// Gets all the users.
			//
			// Responses:
			//	 200: []User Users array.
			//
			r.Get("/", s.getUsers)

			// POST to register
			// swagger:route POST /users auth users registerUser
			//
			// Register an user with email, name and password.
			//
			// The email has to be unique amongst the already registered members.
			//
			// Responses:
			//   201: RegisterResponse The newly created User with an access token.
			//   400: ErrorResponse Invalid payload fields and email already taken error.
			//
			r.Post("/", s.registerUser)


			// Route for logged in current user.
			r.Route("/current", func(r chi.Router) {

				// Validates the Authentication token and injects the user into request context.
				r.Use(s.WithUserAuthenticationCtx)

				// GET for getting current user.
				// swagger:route GET /users/current users currentUser getSignedInUser
				//
				// Get the currently logged in user from Bearer token.
				//
				// Basically a user profile fetcher.
				//
				// Responses:
				//   200: User The currently logged in user.
				//   401: ErrorResponse Authentication error.
				//
				// Security:
				//   bearer_auth:
				//
				r.Get("/", s.getCurrentUser)

				// PATCH for updating user information.
				// swagger:route PATCH /users/current users current_user modifyCurrentUser
				//
				// Modify user details, which are email and the name.
				//
				// The email has to be unique amongst the already registered users. At least one of these parameters have to be supplied.
				//
				// Responses:
				//   200: User User with updated fields.
				//   400: ErrorResponse Invalid payload fields and email already taken error.
				//   401: ErrorResponse Authentication error.
				//
				// Security:
				//   bearer_auth:
				//
				r.Patch("/", s.modifyUser)

				// DELETE for unregistering.
				// swagger:route DELETE /users/current users current_user unregisterCurrentUser
				//
				// Unregister the user that is denoted from the Bearer token.
				//
				// Supply the api with email and current password, so that that the user confirms his choice.
				//
				// Responses:
				//   204: description:User unregistered.
				//   400: ErrorResponse Invalid credentials.
				//   401: ErrorResponse Authentication error.
				//
				// Security:
				//   bearer_auth:
				//
				r.Delete("/", s.unregisterUser)

				// Routes for current user group.
				r.Route("/group", func(r chi.Router) {

					// This will get the group from Datastore and put it into request context.
					r.Use(s.CurrentUserGroupCtx)

					// GET for fetching the current group.
					// swagger:route GET /users/current/group users current_user groups getCurrentUserGroup
					//
					// Get the group from the currently logged in user.
					//
					// Responses:
					//   200: Group Current users group.
					//   401: ErrorResponse Authentication error.
					//   404: description:User doesn't have an assigned group.
					//
					// Security:
					//   bearer_auth:
					//
					r.Get("/", s.getGroup)

					// POST for joining a group.
					// swagger:route POST /users/current/group users current_user groups joinGroup
					//
					// Join a group denoted by group_id.
					//
					// The group should exist and the user should not be in a group already. If these circumstances are not respected an error will be returned.
					//
					// Responses:
					//   201: Group The joined group.
					//   400: ErrorResponse Already in group or group not found error.
					//   401: ErrorResponse Authentication error.
					//
					// Security:
					//   bearer_auth:
					//
					r.Post("/", s.joinGroup)

					// DELETE for leaving the current group.
					// swagger:route DELETE /users/current/group users current_user groups leaveGroup
					//
					// Leave the current group.
					//
					// In any instance the API returns a successful delete response.
					//
					// Responses:
					//   204: description:Successful leave group.
					//   401: ErrorResponse Authentication error.
					//
					// Security:
					//   bearer_auth:
					//
					r.Delete("/", s.leaveGroup)
				})
			})

			// Routes for specific user.
			r.Route("/{" + userIdParam + "}", func(r chi.Router) {
				// Injects the user with user_id into request context.
				r.Use(s.UserCtx)

				// GET a user with user_id.
				// swagger:route GET /users/{user_id} users getUser
				//
				// Get the user which is denoted by user_id.
				//
				// Responses:
				//   200: User User denoted by user_id.
				//   404: description:User with not found.
				//
				r.Get("/", s.getUser)
			})
		})

		// Routes for Groups.
		r.Route("/groups", func(r chi.Router) {

			// GET gets all the groups.
			// swagger:route GET /groups groups getGroups
			//
			// Gets all the groups.
			//
			// Responses:
			//   200: []Group Groups array.
			//
			r.Get("/", s.getGroups)

			// POST to create a group.
			// swagger:route POST /groups groups createGroup
			//
			// Creates a group with a name, it that is not already taken.
			//
			// Responses:
			//   201: Group The created group.
			//   400: description:Invalid payload fields or name already taken.
			//
			r.Post("/", s.createGroup)

			r.Route("/{" + groupIdParam + "}", func(r chi.Router) {

				// Injects the group with group_id into request context.
				r.Use(s.GroupCtx)

				// GET to fetch the group.
				// swagger:route GET /groups/{group_id} groups getGroup
				//
				// Get a group which is denoted by group_id.
				//
				// Responses:
				//   200: Group Group denoted by group_id.
				//   404: description:Group not found.
				//
				r.Get("/", s.getGroup)

				// DELETE to delete the group.
				// swagger:route DELETE /groups/{group_id} groups deleteGroup
				//
				// Delete the group denoted by group_id.
				//
				// Responses:
				//   204: description:Successful delete.
				//
				r.Delete("/", s.deleteGroup)

				// PATCH to modify group name.
				// swagger:route PATCH /groups/{group_id} groups modifyGroup
				//
				// Change the group name.
				//
				// If the name is taken it will return a bad request error with explanation.
				//
				// Responses:
				//   200: Group Group with newly modified fields.
				//   404: description:Group not found.
				//
				r.Patch("/", s.modifyGroup)

				r.Route("/users", func(r chi.Router) {

					// GET all members of a group.
					// swagger:route GET /groups/{group_id}/users groups users getMembersOfGroup
					//
					// Gets all the members of a group denoted by group_id.
					//
					// Responses:
					//   200: []User Members of a the group.
					//   404: description:Group not found.
					//
					r.Get("/", s.getUsersOfGroup)
				})
			})
		})
	})
}