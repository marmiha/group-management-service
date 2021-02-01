package restapi

import "github.com/go-chi/chi"

func (s *Server) setupEndpoints(r *chi.Mux) {
	r.Route("/api/v1", func(r chi.Router) {
		// Routes for Users
		r.Route("/users", func(r chi.Router) {
			// GET gets all the users.
			r.Get("/", s.getUsers)

			// Route for logged in current user.
			r.Route("/current", func(r chi.Router) {
				r.Use(s.WithUserAuthenticationCtx)
				// GET for getting current user.
				// PATCH for updating user information.
				// DELETE for unregistering.
				r.Route("/group", func(r chi.Router) {
					// GET for fetching the current group.
					// POST for joining a group.
					// DELETE for leaving the current group.
				})
			})

			// Route for specific user
			r.Route("/{user_id}", func(r chi.Router) {
				// GET
			})
		})

		// Routes for Groups.
		r.Route("/groups", func(r chi.Router) {
			// GET gets all the groups.
			r.Get("/", s.getGroups)
			// POST to create a group.
			r.Post("/", s.createGroup)

			r.Route("/{"+groupIdParam+"}", func(r chi.Router) {
				r.Use(s.GroupCtx)
				// GET to fetch the group.
				// DELETE to delete the group.
				// PATCH to modify group name.
			})
		})
	}
}