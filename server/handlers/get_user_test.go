package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"net/http/httptest"
	"server/auth"
	"server/db"
	"server/handlers/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUser(t *testing.T) {
	t.Run("successful get user", func(t *testing.T) {
		w, r := httptest.NewRecorder(), httptest.NewRequest("GET", "/user", nil)
		testUserId, testUserIdStr := MockUuid()
		env := &MockedEnv{Queries: new(MockedQueries)}

		expectedUser := db.User{
			ID:    testUserId,
			Email: "test@example.com",
			Alias: "testuser",
		}

		ctx := context.WithValue(r.Context(), auth.AuthContextKey, map[string]string{
			"id": testUserIdStr,
		})
		r = r.WithContext(ctx)

		env.GetQueries().(*MockedQueries).On("GetUser", r.Context(), testUserId).Return(expectedUser, nil)

		err := GetUser(env, w, r)
		assert.NoError(t, err)

		var response db.User
		err = json.NewDecoder(w.Body).Decode(&response)
		assert.NoError(t, err)
		assert.Equal(t, expectedUser, response)
	})

	t.Run("user not found", func(t *testing.T) {
		w, r := httptest.NewRecorder(), httptest.NewRequest("GET", "/user", nil)
		testUserId, testUserIdStr := MockUuid()
		env := &MockedEnv{Queries: new(MockedQueries)}

		ctx := context.WithValue(r.Context(), auth.AuthContextKey, map[string]string{
			"id": testUserIdStr,
		})
		r = r.WithContext(ctx)

		env.GetQueries().(*MockedQueries).On("GetUser", r.Context(), testUserId).Return(db.User{}, errors.New("user not found"))

		err := GetUser(env, w, r)
		assert.Error(t, err)
		assert.Equal(t, 404, err.(utils.StatusError).Code)
	})
}
