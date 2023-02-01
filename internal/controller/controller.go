package controller

import (
	"fmt"

	"github.com/alsolovyev/dummy-api/internal/entity"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func New(f entity.FileUseCaser) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/ping", handlePing)
		r.With(filenameMiddleware).Get(fmt.Sprintf("/file/{%s}", filenameKey), handleFile(f))
	})

	return r
}
