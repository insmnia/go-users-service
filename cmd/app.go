package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/insmnia/go-users-service/database"
	"github.com/insmnia/go-users-service/internal/api"
	"go.uber.org/zap"
)

func InitApp() (app *gin.Engine, err error) {
	app = gin.Default()
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
	return
}
