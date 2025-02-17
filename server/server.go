package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"server/auth"
	"server/handlers"
	handlersUtils "server/handlers/utils"
	"server/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// init auth provider
	auth := auth.InitAuth()

	// init db conn
	dbName := os.Getenv("PGDATABASE")
	dbUser := os.Getenv("PGUSER")
	dbPass := os.Getenv("PGPASSWORD")
	dbHost := os.Getenv("PGHOST")
	dbPort := os.Getenv("PGPORT")

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUser, dbPass, dbHost, dbPort, dbName),
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to DB")
	}

	// Migrate the schema
	db.AutoMigrate(&models.User{})

	// init server env
	env := &handlersUtils.Env{DB: db, Auth: auth}

	// authenticated routes
	authMux := http.NewServeMux()
	authMux.Handle("GET /user", handlersUtils.Handler{Env: env, H: handlers.GetUser})
	authMux.Handle("POST /user", handlersUtils.Handler{Env: env, H: handlers.CreateUser})

	mux := http.NewServeMux()
	mux.Handle("/auth/", http.StripPrefix("/auth", auth.Authenticate(authMux)))

	slog.Info("Server starting on port 8080")
	log.Fatal(http.ListenAndServe(":8080", handlersUtils.CorsHandler(mux)))
}
