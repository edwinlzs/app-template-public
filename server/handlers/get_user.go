package handlers

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"server/handlers/utils"
	"server/models"
)

// Retrieves user data
func GetUser(env utils.ServerEnv, w http.ResponseWriter, r *http.Request) error {
	slog.Info("Getting user")

	userContext, err := utils.GetUserContext(r)
	if err != nil {
		return err
	}

	userId, err := utils.GetUserId(userContext)
	if err != nil {
		return err
	}

	var user models.User
	result := env.GetDB().First(&user, userId)
	if result.RowsAffected == 0 {
		slog.Warn(fmt.Sprintf("User not found: %s", err))
		return utils.StatusError{Code: 404, Err: errors.New("failed to get user")}
	}

	slog.Info(fmt.Sprintf("User found: %s", user.Email))

	utils.JSONResponse(w, user)
	return nil
}
