package rest

import (
	"github.com/go-chi/chi"

	"github.com/cortezaproject/corteza-server/federation/rest/handlers"
)

func MountRoutes(r chi.Router) {
	var (
		module = Module{}.New()
		record = Record{}.New()
	)

	// Protect all _private_ routes
	r.Group(func(r chi.Router) {
		// r.Use(auth.MiddlewareValidOnly)
		// r.Use(middlewareAllowedAccess)

		handlers.NewModule(module).MountRoutes(r)
		handlers.NewRecord(record).MountRoutes(r)
	})
}
