package utils

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCorsHandler(t *testing.T) {
	t.Run("adds CORS headers", func(t *testing.T) {
		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})
		corsHandler := CorsHandler(handler)

		w := httptest.NewRecorder()
		r := httptest.NewRequest("GET", "/", nil)

		corsHandler.ServeHTTP(w, r)

		assert.Equal(t, "*", w.Header().Get("Access-Control-Allow-Origin"))
		assert.Equal(t, "POST, GET, OPTIONS, PUT, DELETE", w.Header().Get("Access-Control-Allow-Methods"))
	})

	t.Run("handles OPTIONS request", func(t *testing.T) {
		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			t.Error("Handler should not be called for OPTIONS")
		})
		corsHandler := CorsHandler(handler)

		w := httptest.NewRecorder()
		r := httptest.NewRequest("OPTIONS", "/", nil)

		corsHandler.ServeHTTP(w, r)
	})
}
