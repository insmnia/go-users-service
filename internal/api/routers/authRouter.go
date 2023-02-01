package routers

import (
	"github.com/gin-gonic/gin"
	r "github.com/insmnia/go-users-service/internal/repository"
	"github.com/insmnia/go-users-service/internal/service"
	"go.uber.org/zap"
)

type AuthRoutesController struct {
	authService *service.AuthService
}

func NewAuthRoutes(repo *r.UserRepository, logger *zap.SugaredLogger) *AuthRoutesController {
	return &AuthRoutesController{authService: service.NewAuthService(repo, logger)}
}

func (controller *AuthRoutesController) InitAuthRoutes(apiGroup gin.IRoutes) {
	apiGroup.POST("/register", controller.SignUpUser)
	apiGroup.POST("/login", controller.SignInUser)
	apiGroup.POST("/refresh", controller.ExpandUserToken)
}

func (controller *AuthRoutesController) SignUpUser(ctx *gin.Context) {
	controller.authService.Create(ctx)
}

func (controller *AuthRoutesController) SignInUser(ctx *gin.Context) {
	controller.authService.Authorize(ctx)
}

func (controller *AuthRoutesController) ExpandUserToken(ctx *gin.Context) {
	controller.authService.RefreshToken(ctx)
}
