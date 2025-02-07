package handlers

import (
	"context"
	"server/db"
	"server/handlers/utils"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/mock"
)

type MockedQueries struct {
	mock.Mock
}

func (m *MockedQueries) GetUser(ctx context.Context, id pgtype.UUID) (db.User, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(db.User), args.Error(1)
}

func (m *MockedQueries) CreateUser(ctx context.Context, params db.CreateUserParams) (db.User, error) {
	args := m.Called(ctx, params)
	return args.Get(0).(db.User), args.Error(1)
}

type MockedEnv struct {
	Queries *MockedQueries
}

func (m *MockedEnv) GetQueries() utils.DbQueries {
	return m.Queries
}

// Creates a Postgres UUID and its string counterpart
func MockUuid() (pgtype.UUID, string) {
	id := uuid.New().String()
	var pgId pgtype.UUID
	(pgId).Scan(id)
	return pgId, id
}
