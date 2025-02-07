package utils

import (
	"context"
	"server/auth"
	"server/db"

	"github.com/jackc/pgx/v5/pgtype"
)

type ServerEnv interface {
	GetQueries() DbQueries
}

type DbQueries interface {
	GetUser(ctx context.Context, id pgtype.UUID) (db.User, error)
	CreateUser(ctx context.Context, params db.CreateUserParams) (db.User, error)
}

// App-wide server environment
type Env struct {
	Queries *db.Queries
	auth.Auth
}

func (e *Env) GetQueries() DbQueries {
	return e.Queries
}
