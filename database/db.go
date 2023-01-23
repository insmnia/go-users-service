package database

import (
	"github.com/insmnia/go-users-service/core/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	dbConfig, err := config.LoadDatabaseConfig(".")
	if err != nil {
		log.Fatalf("Couldn't read database config due to %s", err.Error())
	}
	db, err := gorm.Open("postgres", dbConfig.ToConnectionString())
	if err != nil {
		log.Fatalf("Couldn't connect to database due to %s", err.Error())
	}
	DB = db
	return DB
}

func GetDB() *gorm.DB {
	return DB
}
