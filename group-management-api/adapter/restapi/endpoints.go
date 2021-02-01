package restapi

import "github.com/go-chi/chi"

func (s *Server) setupEndpoints(r *chi.Mux) {
	r.Route("/api/v1", func(r chi.Router) {

		// For logging the user in.
		r.Route("/login", func(r chi.Router) {
			r.Post("/", s.loginUser)
		})

		// Routes for Users
		r.Route("/users", func(r chi.Router) {
			// GET gets all the users.
			r.Get("/", s.getUsers)
			// POST to register
			r.Post("/", s.registerUser)


			// Route for logged in current user.
			r.Route("/current", func(r chi.Router) {
				r.Use(s.WithUserAuthenticationCtx)
				// GET for getting current user.
				r.Get("/", s.getCurrentUser)
				// PATCH for updating user information.
				r.Patch("/", s.modifyUser)
				// DELETE for unregistering.
				r.Delete("/", s.unregisterUser)

				r.Route("/group", func(r chi.Router) {
					r.Use(s.CurrentUserGroupCtx)
					// GET for fetching the current group.
					r.Get("/", s.getGroup)
					// POST for joining a group.
					r.Post("/", s.joinGroup)
					// DELETE for leaving the current group.
					r.Delete("/", s.leaveGroup)
				})
			})

			// Route for specific user
			r.Route("/{" + userIdParam + "}", func(r chi.Router) {
				r.Use(s.UserCtx)
				// GET
				r.Get("/", s.getUser)
			})
		})

		// Routes for Groups.
		r.Route("/groups", func(r chi.Router) {
			// GET gets all the groups.
			r.Get("/", s.getGroups)
			// POST to create a group.
			r.Post("/", s.createGroup)

			r.Route("/{" + groupIdParam + "}", func(r chi.Router) {
				r.Use(s.GroupCtx)
				// GET to fetch the group.
				r.Get("/", s.getGroup)
				// DELETE to delete the group.
				r.Delete("/", s.deleteGroup)
				// PATCH to modify group name.
				r.Patch("/", s.modifyGroup)
			})
		})
	})
}