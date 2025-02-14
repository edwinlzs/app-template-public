package utils

import (
	"server/auth"

	"gorm.io/gorm"
)

type ServerEnv interface {
	GetDB() *gorm.DB
}

// App-wide server environment
type Env struct {
	*gorm.DB
	auth.Auth
}

func (e *Env) GetDB() *gorm.DB {
	return e.DB
}
