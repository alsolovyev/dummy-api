package controller

import (
	"fmt"
	"net/http"
	"strconv"
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
	r.Use(statusCodeMiddleware)

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

// A middleware that can be used to manipulate the HTTP response status code.
// It inspects the query parameters of the URL in the request. If the
// "status_code" query parameter is present and its value is a valid HTTP
// status code between 100 and 599, the middleware sets the status code
// of the HTTP response to the specified value.
func statusCodeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()

		if q["status_code"] != nil {
			d, err := strconv.Atoi(q["status_code"][0])

			if err == nil && d > 99 && d <= 599 {
				w.WriteHeader(d)
			}
		}

		next.ServeHTTP(w, r)
	})
}
