package controller

import (
	"github.com/alsolovyev/dummy-api/internal/entity"
	"github.com/go-chi/chi/v5"
)

func New(f entity.FileUseCaser) *chi.Mux {
	r := chi.NewRouter()

	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/ping", handlePing)
		r.With(filenameMiddleware).Get("/file/{name}", handleFile(f))
	})

	return r
}
