package auth

import (
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/golang-jwt/jwt/v5"
)

func TestGetAuthTokenFromRequest(t *testing.T) {
	req, err := http.NewRequest("GET", "/user", nil)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Authorization", "Bearer token")

	token := GetAuthTokenFromRequest(req)
	if token != "token" {
		t.Errorf("Expected token to be 'token', got %s", token)
	}
}

func TestVerifyToken(t *testing.T) {
	auth := Auth{"test-secret"}

	claims, err := auth.VerifyToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiZW1haWwiOiJ0ZXN0QGVtYWlsLmNvbSIsImlhdCI6MTUxNjIzOTAyMn0.2r_X2IDzy1Vn78Tx4uMCE-_NdKNcKEK12b5Tb4NOsfA")
	if err != nil {
		t.Fatal(err)
	}

	expected := jwt.MapClaims{
		"sub":   "1234567890",
		"name":  "John Doe",
		"email": "test@email.com",
		"iat":   float64(1516239022),
	}

	if !reflect.DeepEqual(claims, expected) {
		t.Error("Expected claims to be", expected, "got", claims)
	}
}

func TestAuthenticate(t *testing.T) {
	log.SetOutput(io.Discard) // discard logs from tested function

	auth := Auth{"test-secret"}

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	authHandler := auth.Authenticate(handler)

	req, err := http.NewRequest("GET", "/user", nil)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiZW1haWwiOiJ0ZXN0QGVtYWlsLmNvbSIsImlhdCI6MTUxNjIzOTAyMn0.2r_X2IDzy1Vn78Tx4uMCE-_NdKNcKEK12b5Tb4NOsfA")

	rec := httptest.NewRecorder()
	authHandler.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("Expected status code to be %d, got %d", http.StatusOK, rec.Code)
	}
}
