package controller

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/alsolovyev/dummy-api/internal/entity"
	"github.com/go-chi/chi/v5"
)

const filenameKey = "name"

func filenameMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		f := chi.URLParam(r, filenameKey)

		ctx := context.WithValue(r.Context(), filenameKey, f)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}


func handleFile(f entity.FileUseCaser) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		p := r.Context().Value(filenameKey).(string)

		d, err := f.GetFile(p)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(d)
	}
}
