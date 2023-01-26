package main

import (
	"github.com/gin-gonic/gin"
	"github.com/insmnia/go-users-service/api"
	"github.com/insmnia/go-users-service/database"
	"go.uber.org/zap"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	app := gin.Default()
	db, err := database.InitDB()
	if err != nil {
		return
	}
	database.MigrateModels(db)

	zapLogger, _ := zap.NewProduction()
	defer func(logger *zap.Logger) {
		err := logger.Sync()
		if err != nil {
			return
		}
	}(zapLogger)
	logger := zapLogger.Sugar()

	api.SetUpRoutes(app, db, logger)
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
