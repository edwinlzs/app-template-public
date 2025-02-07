package utils

import (
	"log"
	"net/http"
)

type Handler struct {
	Env ServerEnv
	H   func(e ServerEnv, w http.ResponseWriter, r *http.Request) error
}

// allows Handler type to satisfy http.Handler
func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := h.H(h.Env, w, r)
	if err != nil {
		switch e := err.(type) {
		case Error:
			log.Printf("HTTP %d - %s", e.Status(), e)
			http.Error(w, e.Error(), e.Status())
		default:
			http.Error(w, http.StatusText(http.StatusInternalServerError),
				http.StatusInternalServerError)
		}
	}
}
