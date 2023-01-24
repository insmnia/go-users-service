package repository

import (
	"github.com/google/uuid"
	"github.com/insmnia/go-users-service/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (repo *UserRepository) Create(u models.User) (uuid.UUID, error) {
	u.UUID = uuid.New()
	result := repo.db.Create(&u)
	return u.UUID, result.Error
}
