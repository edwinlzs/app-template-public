package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID    uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()"`
	Email string
	Alias string
	gorm.Model
}
