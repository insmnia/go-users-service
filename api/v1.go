package api

import (
	"github.com/gin-gonic/gin"
	"github.com/insmnia/go-users-service/api/routers"
	"github.com/insmnia/go-users-service/core/middlewares"
	"github.com/insmnia/go-users-service/repository"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func SetUpRoutes(app *gin.Engine, db *gorm.DB, logger *zap.SugaredLogger) {
	userRepo := repository.NewUserRepository(db)
	userRoutes := routers.NewUserRoutes(userRepo, logger)
	authRoutes := routers.NewAuthRoutes(userRepo, logger)
	api := app.Group("/api")
	userRoutes.InitUserRoutes(api.Group("/users").Use(middlewares.RequiresAuth()))
	authRoutes.InitAuthRoutes(api.Group("/auth"))
}
