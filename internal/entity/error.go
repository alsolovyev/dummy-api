package entity

import (
	"encoding/json"
	"net/http"
)

type Error struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func NewError(s int, m string) *Error {
	return &Error{
		Message: m,
		Status:  s,
	}
}

func (e *Error) Error() string {
	return e.Message
}

func (e *Error) Render(w http.ResponseWriter) error {
	w.WriteHeader(e.Status)

	return json.NewEncoder(w).Encode(e)
}
