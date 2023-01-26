package database

import (
	"github.com/insmnia/go-users-service/core/config"
	"github.com/insmnia/go-users-service/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func InitDB() (*gorm.DB, error) {
	dbConfig, err := config.LoadDatabaseConfig("./env")
	if err != nil {
		log.Fatalf("Couldn't read database config due to %s", err.Error())
		return nil, err
	}
	db, err := gorm.Open(postgres.Open(dbConfig.ToConnectionString()), &gorm.Config{})
	if err != nil {
		log.Fatalf("Couldn't connect to database due to %s", err.Error())
		return nil, err
	}
	return db, nil
}

func MigrateModels(db *gorm.DB) {
	err := db.AutoMigrate(&models.User{})
	if err != nil {
		return
	}
}
