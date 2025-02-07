package utils

import (
	"context"
	"net/http/httptest"
	"server/auth"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUserContext(t *testing.T) {
	t.Run("successful extraction", func(t *testing.T) {
		r := httptest.NewRequest("GET", "/", nil)
		userData := map[string]string{"id": "123", "email": "test@example.com"}
		ctx := context.WithValue(r.Context(), auth.AuthContextKey, userData)
		r = r.WithContext(ctx)

		result, err := GetUserContext(r)
		assert.NoError(t, err)
		assert.Equal(t, userData, result)
	})

	t.Run("missing context", func(t *testing.T) {
		r := httptest.NewRequest("GET", "/", nil)
		_, err := GetUserContext(r)
		assert.Error(t, err)
		assert.Equal(t, 400, err.(StatusError).Code)
	})
}

func TestGetUserId(t *testing.T) {
	t.Run("valid UUID", func(t *testing.T) {
		userContext := map[string]string{"id": "123e4567-e89b-12d3-a456-426614174000"}
		_, err := GetUserId(userContext)
		assert.NoError(t, err)
	})

	t.Run("invalid UUID", func(t *testing.T) {
		userContext := map[string]string{"id": "invalid-uuid"}
		_, err := GetUserId(userContext)
		assert.Error(t, err)
		assert.Equal(t, 400, err.(StatusError).Code)
	})
}
