package repository

import (
	"github.com/google/uuid"
	"github.com/insmnia/go-users-service/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

type IUserRepository interface {
	Create(u models.User) (models.User, error)
	GetById(id uuid.UUID) (models.User, error)
	GetByName(name string) (models.User, error)
	Update(u models.UpdateUser) (models.User, error)
	Delete(id uuid.UUID) error
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (repo *UserRepository) Create(u models.User) (models.User, error) {
	dbUser := models.User{Username: u.Username, Password: u.Password}
	result := repo.db.Create(&dbUser)
	return dbUser, result.Error
}

func (repo *UserRepository) Update(id uuid.UUID, u models.UpdateUser) (models.User, error) {
	dbUser, _ := repo.GetById(id)
	result := repo.db.Model(&dbUser).Updates(models.User{Username: u.Username})
	return dbUser, result.Error
}

func (repo *UserRepository) Delete(id uuid.UUID) error {
	result := repo.db.Where("UUID = ?", id).Delete(&models.User{})
	return result.Error
}

func (repo *UserRepository) GetById(id uuid.UUID) (models.User, error) {
	var user models.User
	result := repo.db.Where(models.User{UUID: id}).First(&user)
	return user, result.Error
}

func (repo *UserRepository) GetForceByName(name string) (models.User, error) {
	var user models.User
	result := repo.db.Unscoped().Where(models.User{Username: name}).First(&user)
	return user, result.Error
}
