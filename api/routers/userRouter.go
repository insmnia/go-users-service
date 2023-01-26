package routers

import (
	"github.com/gin-gonic/gin"
	r "github.com/insmnia/go-users-service/repository"
	"github.com/insmnia/go-users-service/service"
	"go.uber.org/zap"
)

type UserRoutesController struct {
	userService *service.UserService
}

func NewUserRoutes(repo *r.UserRepository, logger *zap.SugaredLogger) *UserRoutesController {
	return &UserRoutesController{userService: service.NewUserService(repo, logger)}
}

func (controller *UserRoutesController) InitUserRoutes(apiGroup gin.IRoutes) {
	apiGroup.GET("/:uuid", controller.GetUser)
	apiGroup.PATCH("/:uuid", controller.UpdateUser)
	apiGroup.DELETE("/:uuid", controller.DeleteUser)
}

func (controller *UserRoutesController) GetUser(ctx *gin.Context) {
	controller.userService.Get(ctx)
}

func (controller *UserRoutesController) DeleteUser(ctx *gin.Context) {
	controller.userService.Delete(ctx)
}

func (controller *UserRoutesController) UpdateUser(ctx *gin.Context) {
	controller.userService.Update(ctx)
}
