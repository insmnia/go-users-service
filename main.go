package main

import (
	"github.com/gin-gonic/gin"
	"github.com/insmnia/go-users-service/database"
	"github.com/insmnia/go-users-service/models"
	"github.com/insmnia/go-users-service/repository"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	app := gin.Default()
	database.InitDB()
	database.MigrateModels()
	app.GET("/users", func(ctx *gin.Context) {
		repo := repository.NewUserRepository(database.GetDB())
		id, err := repo.Create(models.User{Username: "123", Password: "123"})
		if err != nil {
			return
		}
		ctx.JSON(201, gin.H{"id": id})
	})
	log.Print("Server started")
	go func() {
		err := app.Run()
		if err != nil {
			log.Fatalf("Failed to start server due to %s", err.Error())
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	log.Print("Shutting down server")
}
