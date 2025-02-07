package auth

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

type contextKey string

const AuthContextKey contextKey = "auth"

func InitAuth() Auth {
	jwtSecret := os.Getenv("JWT_SECRET")
	slog.Info("Initialized auth")

	return Auth{jwtSecret}
}

type Auth struct {
	jwtSecret string
}

func GetAuthTokenFromRequest(r *http.Request) string {
	authHeaderParts := strings.Split(r.Header.Get("Authorization"), "Bearer ")

	if len(authHeaderParts) != 2 {
		return ""
	}

	return authHeaderParts[1]
}

func (auth *Auth) VerifyToken(tokenStr string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(auth.jwtSecret), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		return claims, nil
	}

	return nil, fmt.Errorf("no claims found")
}

func (auth *Auth) Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Authenticating")
		tokenStr := GetAuthTokenFromRequest(r)
		claims, err := auth.VerifyToken(tokenStr)
		if err != nil {
			log.Printf("HTTP %d - %s", 400, err)
			http.Error(w, err.Error(), 400)
			return
		}
		slog.Info("Authenticated")

		user := make(map[string]string)
		user["id"] = claims["sub"].(string)
		user["email"] = claims["email"].(string)
		r = r.WithContext(context.WithValue(r.Context(), AuthContextKey, user))

		next.ServeHTTP(w, r)
	})
}
