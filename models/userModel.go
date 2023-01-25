package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UUID     uuid.UUID `json:"uuid" gorm:"type:uuid;primary_key;primaryKey"`
	Username string    `json:"username" gorm:"uniqueIndex"`
	Password string    `json:"password" gorm:"not null"`
}

type CreateUser struct {
	Username string
	Password string
}
type UpdateUser struct {
	Username string
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	u.UUID = uuid.New()
	return nil
}
