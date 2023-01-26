package api

import (
	"github.com/gin-gonic/gin"
	routers2 "github.com/insmnia/go-users-service/internal/api/routers"
	"github.com/insmnia/go-users-service/internal/middlewares"
	"github.com/insmnia/go-users-service/internal/repository"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func SetUpRoutes(app *gin.Engine, db *gorm.DB, logger *zap.SugaredLogger) {
	userRepo := repository.NewUserRepository(db)
	userRoutes := routers2.NewUserRoutes(userRepo, logger)
	authRoutes := routers2.NewAuthRoutes(userRepo, logger)
	api := app.Group("/api")
	userRoutes.InitUserRoutes(api.Group("/users").Use(middlewares.RequiresAuth()))
	authRoutes.InitAuthRoutes(api.Group("/auth"))
}
