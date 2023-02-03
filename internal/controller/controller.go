package controller

import (
	"fmt"
	"net/http"
	"time"

	"github.com/alsolovyev/dummy-api/internal/entity"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func New(f entity.FileUseCaser) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(delayMiddleware)

	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/ping", handlePing)
		r.With(filenameMiddleware).Get(fmt.Sprintf("/file/{%s}", filenameKey), handleFile(f))
	})

	return r
}

func delayMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()

		if q["delay"] != nil {
			d, err := time.ParseDuration(q["delay"][0])

			if err != nil {
				entity.NewError(http.StatusBadRequest, "Invalid dealy value").Render(w)
				return
			}
			time.Sleep(d)
		}

		next.ServeHTTP(w, r)
	})
}
