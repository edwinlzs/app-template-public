package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"server/auth"
	"server/db"
	"server/handlers"
	handlersUtils "server/handlers/utils"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
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
	dbConnUri := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, dbConnUri)
	if err != nil {
		log.Fatal("Error connecting to DB")
	}
	defer conn.Close(context.Background())
	queries := db.New(conn)

	// init server env
	env := &handlersUtils.Env{Queries: queries, Auth: auth}

	// authenticated routes
	authMux := http.NewServeMux()
	authMux.Handle("GET /user", handlersUtils.Handler{Env: env, H: handlers.GetUser})
	authMux.Handle("POST /user", handlersUtils.Handler{Env: env, H: handlers.CreateUser})

	mux := http.NewServeMux()
	mux.Handle("/auth/", http.StripPrefix("/auth", auth.Authenticate(authMux)))

	slog.Info("Server starting on port 8080")
	log.Fatal(http.ListenAndServe(":8080", handlersUtils.CorsHandler(mux)))
}
