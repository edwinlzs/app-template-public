package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"server/db"
	"server/handlers/utils"
)

type CreateUserPayload struct {
	Alias string `json:"alias" validate:"required"`
}

// Retrieves user data and registers user if new
func CreateUser(env utils.ServerEnv, w http.ResponseWriter, r *http.Request) error {
	slog.Info("Registering user")
	var payload CreateUserPayload

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		return utils.StatusError{Code: 400, Err: err}
	}

	userContext, err := utils.GetUserContext(r)
	if err != nil {
		return utils.StatusError{Code: 403, Err: err}
	}

	userId, err := utils.GetUserId(userContext)
	if err != nil {
		return utils.StatusError{Code: 403, Err: err}
	}

	alias, email := payload.Alias, userContext["email"]
	if alias == "" || email == "" {
		return utils.StatusError{Code: 400, Err: errors.New("empty user alias or email")}
	}

	ctx := r.Context()
	user, err := env.GetQueries().CreateUser(ctx, db.CreateUserParams{ID: userId, Alias: alias, Email: email})
	if err != nil {
		slog.Warn(fmt.Sprintf("Failed to create user: %s", err))
		return utils.StatusError{Code: 500, Err: errors.New("failed to create user")}
	}

	slog.Info(fmt.Sprintf("User with email '%s' registered as '%s'", email, alias))
	utils.JSONResponse(w, user)
	return nil
}
