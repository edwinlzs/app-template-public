package handlers

import (
	"bytes"
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

func TestCreateUser(t *testing.T) {
	t.Run("successful create user", func(t *testing.T) {
		payload := CreateUserPayload{Alias: "testuser"}
		payloadBytes, _ := json.Marshal(payload)
		w, r := httptest.NewRecorder(), httptest.NewRequest("POST", "/user", bytes.NewBuffer(payloadBytes))
		testUserId, testUserIdStr := MockUuid()
		env := &MockedEnv{Queries: new(MockedQueries)}

		expectedUser := db.User{
			ID:    testUserId,
			Email: "test@example.com",
			Alias: "testuser",
		}

		ctx := context.WithValue(r.Context(), auth.AuthContextKey, map[string]string{
			"id":    testUserIdStr,
			"email": "test@example.com",
		})
		r = r.WithContext(ctx)

		env.GetQueries().(*MockedQueries).On("CreateUser", r.Context(), db.CreateUserParams{ID: testUserId, Alias: "testuser", Email: "test@example.com"}).Return(expectedUser, nil)

		err := CreateUser(env, w, r)
		assert.NoError(t, err)

		var response db.User
		err = json.NewDecoder(w.Body).Decode(&response)
		assert.NoError(t, err)
		assert.Equal(t, expectedUser, response)
	})

	t.Run("empty alias or email", func(t *testing.T) {
		payload := CreateUserPayload{Alias: ""}
		payloadBytes, _ := json.Marshal(payload)
		w, r := httptest.NewRecorder(), httptest.NewRequest("POST", "/user", bytes.NewBuffer(payloadBytes))
		_, testUserIdStr := MockUuid()
		env := &MockedEnv{Queries: new(MockedQueries)}

		ctx := context.WithValue(r.Context(), auth.AuthContextKey, map[string]string{
			"id":    testUserIdStr,
			"email": "",
		})
		r = r.WithContext(ctx)

		err := CreateUser(env, w, r)
		assert.Error(t, err)
		assert.Equal(t, 400, err.(utils.StatusError).Code)
	})

	t.Run("failed to decode payload", func(t *testing.T) {
		w, r := httptest.NewRecorder(), httptest.NewRequest("POST", "/user", bytes.NewBuffer([]byte("invalid payload")))
		env := &MockedEnv{Queries: new(MockedQueries)}

		err := CreateUser(env, w, r)
		assert.Error(t, err)
		assert.Equal(t, 400, err.(utils.StatusError).Code)
	})

	t.Run("failed to create user", func(t *testing.T) {
		payload := CreateUserPayload{Alias: "testuser"}
		payloadBytes, _ := json.Marshal(payload)
		w, r := httptest.NewRecorder(), httptest.NewRequest("POST", "/user", bytes.NewBuffer(payloadBytes))
		testUserId, testUserIdStr := MockUuid()
		env := &MockedEnv{Queries: new(MockedQueries)}

		ctx := context.WithValue(r.Context(), auth.AuthContextKey, map[string]string{
			"id":    testUserIdStr,
			"email": "test@example.com",
		})
		r = r.WithContext(ctx)

		env.GetQueries().(*MockedQueries).On("CreateUser", r.Context(), db.CreateUserParams{ID: testUserId, Alias: "testuser", Email: "test@example.com"}).Return(db.User{}, errors.New("failed to create user"))

		err := CreateUser(env, w, r)
		assert.Error(t, err)
		assert.Equal(t, 500, err.(utils.StatusError).Code)
	})
}
