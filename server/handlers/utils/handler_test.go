package utils

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

type mockEnv struct{}

func (m *mockEnv) GetDB() *gorm.DB {
	return nil
}

func TestHandler(t *testing.T) {
	env := &mockEnv{}

	t.Run("successful handler", func(t *testing.T) {
		handler := Handler{
			Env: env,
			H: func(e ServerEnv, w http.ResponseWriter, r *http.Request) error {
				w.WriteHeader(http.StatusOK)
				return nil
			},
		}

		w := httptest.NewRecorder()
		r := httptest.NewRequest("GET", "/", nil)

		handler.ServeHTTP(w, r)
		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("handles StatusError", func(t *testing.T) {
		handler := Handler{
			Env: env,
			H: func(e ServerEnv, w http.ResponseWriter, r *http.Request) error {
				return StatusError{Code: http.StatusBadRequest, Err: errors.New("bad request")}
			},
		}

		w := httptest.NewRecorder()
		r := httptest.NewRequest("GET", "/", nil)

		handler.ServeHTTP(w, r)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}
