package utils

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"server/auth"

	"github.com/google/uuid"
)

func GetUserContext(r *http.Request) (map[string]string, error) {
	user, ok := r.Context().Value(auth.AuthContextKey).(map[string]string)
	if !ok {
		slog.Warn("Unable to extract user data")
		return user, StatusError{Code: 400, Err: errors.New("unable to extract user data")}
	}

	return user, nil
}

// Extracts user ID from request context
func GetUserId(userContext map[string]string) (uuid.UUID, error) {
	var userId uuid.UUID
	err := (userId).Scan(userContext["id"])
	if err != nil {
		slog.Warn(fmt.Sprintf("UserId: %s is not a valid UUID", userContext["id"]))
		return userId, StatusError{Code: 400, Err: errors.New("unable to extract user id")}
	}

	return userId, nil
}
