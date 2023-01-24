package models

import (
	"github.com/google/uuid"
)

type User struct {
	UUID      uuid.UUID `json:"uuid" gorm:"type:uuid;primary_key;primaryKey"`
	Username  string    `json:"username" gorm:"unique_index"`
	Password  string    `json:"password" gorm:"not null"`
	IsDeleted bool      `json:"isDeleted"`
}

// TODO: fix trigger
// BeforeCreate `Before insert` trigger analogue
//func (u *User) BeforeCreate(tx *gorm.DB) error {
//	u.UUID = uuid.New()
//	return nil
//}
