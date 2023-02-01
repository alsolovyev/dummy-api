package controller

import "github.com/go-chi/chi/v5"

func New() *chi.Mux {
	r := chi.NewRouter()

	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/ping", handlePing)
	})

	return r
}
