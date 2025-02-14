package handlers

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/mock"
)

type MockedDB struct {
	mock.Mock
}

func (m *MockedDB) GetUser(ctx context.Context, id pgtype.UUID) (db.User, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(db.User), args.Error(1)
}

func (m *MockedDB) CreateUser(ctx context.Context, params db.CreateUserParams) (db.User, error) {
	args := m.Called(ctx, params)
	return args.Get(0).(db.User), args.Error(1)
}

type MockedEnv struct {
	DB *MockedDB
}

func (m *MockedEnv) GetDB() *MockedDB {
	return m.DB
}
