package database

import (
	"github.com/insmnia/go-users-service/core/config"
	"github.com/insmnia/go-users-service/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	dbConfig, err := config.LoadDatabaseConfig(".")
	if err != nil {
		log.Fatalf("Couldn't read database config due to %s", err.Error())
	}
	db, err := gorm.Open(postgres.Open(dbConfig.ToConnectionString()), &gorm.Config{})
	if err != nil {
		log.Fatalf("Couldn't connect to database due to %s", err.Error())
	}
	DB = db
	return DB
}

func GetDB() *gorm.DB {
	return DB
}

func MigrateModels() {
	err := DB.AutoMigrate(&models.User{})
	if err != nil {
		return
	}
}
