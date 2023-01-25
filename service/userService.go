package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/insmnia/go-users-service/api/dto"
	"github.com/insmnia/go-users-service/core/errors"
	"github.com/insmnia/go-users-service/core/validators"
	"github.com/insmnia/go-users-service/models"
	"github.com/insmnia/go-users-service/repository"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type UserService struct {
	Repo   *repository.UserRepository
	logger *zap.SugaredLogger
}

func NewUserService(userRepository *repository.UserRepository, logger *zap.SugaredLogger) *UserService {
	return &UserService{userRepository, logger}
}

func (service *UserService) GenerateHashPassword(password string) (string, error) {
	// TODO: move cost to env variable
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(hashedPassword), err
}

func (service *UserService) CheckHashPassword(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err != nil
}

func (service *UserService) Create(ctx *gin.Context) {
	var input models.CreateUser

	if err := ctx.ShouldBindJSON(&input); err != nil {
		service.logger.Errorf("Error while parsing body %s", err)
		bodyError := errors.ParseBodyError{Msg: "Couldn't parse body"}
		bodyError.Raise(ctx)
		return
	}
	if errs := validators.ValidateStruct(input); errs != nil {
		service.logger.Errorf("Couldn't validate input structure: %v", errs)
		validationError := errors.ValidateBodyError{Errors: errs}
		validationError.Raise(ctx)
		return
	}
	if _, err := service.Repo.GetForceByName(input.Username); err == nil {
		ctx.JSON(http.StatusBadRequest, "User with such username already exists!")
		return
	}
	hashedPassword, err := service.GenerateHashPassword(input.Password)
	if err != nil {
		service.logger.Errorf("Couldn't generate password due to: %s", err)
		ctx.Status(http.StatusInternalServerError)
		return
	}

	user, err := service.Repo.Create(models.User{Username: input.Username, Password: hashedPassword})
	if err != nil {
		service.logger.Errorf("Couldn't create user due to: %s", err)
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.JSON(201, dto.UserResponse{
		ID:        user.UUID,
		Username:  user.Username,
		CreatedAt: user.CreatedAt,
	})
}

func (service *UserService) Get(ctx *gin.Context) {
	paramId := ctx.Param("uuid")
	userId, err := uuid.Parse(paramId)
	if err != nil {
		service.logger.Errorf("Couldn't convert %s to UUID due to %s", paramId, err)
		ctx.Status(http.StatusNotFound)
		return
	}
	dbUser, err := service.Repo.GetById(userId)
	if err != nil {
		ctx.Status(http.StatusNotFound)
		return
	}
	ctx.JSON(http.StatusOK, dto.UserResponse{
		ID:        dbUser.UUID,
		Username:  dbUser.Username,
		CreatedAt: dbUser.CreatedAt,
	})
}

func (service *UserService) Delete(ctx *gin.Context) {
	userId, err := uuid.Parse(ctx.Param("uuid"))
	if err != nil {
		service.logger.Error("Cannot parse user id from url")
		ctx.Status(http.StatusNotFound)
		return
	}
	if dbErr := service.Repo.Delete(userId); dbErr != nil {
		service.logger.Errorf("Couldn't delete user with id %s due to %s", userId, dbErr)
		ctx.JSON(http.StatusNotFound, fmt.Sprintf("User with such id is not found"))
		return
	}
	ctx.Status(200)
}

func (service *UserService) Update(ctx *gin.Context) {
	var input models.UpdateUser

	if err := ctx.ShouldBindJSON(&input); err != nil {
		service.logger.Errorf("Error while parsing body %s", err)
		bodyError := errors.ParseBodyError{Msg: "Couldn't parse body"}
		bodyError.Raise(ctx)
		return
	}
	if errs := validators.ValidateStruct(input); errs != nil {
		service.logger.Errorf("Couldn't validate input structure: %v", errs)
		validationError := errors.ValidateBodyError{Errors: errs}
		validationError.Raise(ctx)
		return
	}
	if _, err := service.Repo.GetForceByName(input.Username); err == nil {
		ctx.JSON(http.StatusBadRequest, "User with such username already exists!")
		return
	}
	userId, err := uuid.Parse(ctx.Param("uuid"))
	if err != nil {
		service.logger.Error("Cannot parse user id from url")
		ctx.Status(http.StatusNotFound)
		return
	}
	user, err := service.Repo.Update(userId, input)
	if err != nil {
		service.logger.Errorf("Couldn't update user due to: %s", err)
		ctx.Status(http.StatusInternalServerError)
		return
	}
	ctx.JSON(200, dto.UserResponse{
		ID:        user.UUID,
		Username:  user.Username,
		CreatedAt: user.CreatedAt,
	})
}
