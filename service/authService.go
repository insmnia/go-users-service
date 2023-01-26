package service

import (
	"github.com/gin-gonic/gin"
	"github.com/insmnia/go-users-service/api/dto"
	"github.com/insmnia/go-users-service/core/errors"
	"github.com/insmnia/go-users-service/core/utils"
	"github.com/insmnia/go-users-service/core/validators"
	"github.com/insmnia/go-users-service/models"
	"github.com/insmnia/go-users-service/repository"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type AuthService struct {
	Repo   *repository.UserRepository
	logger *zap.SugaredLogger
}

func NewAuthService(r *repository.UserRepository, l *zap.SugaredLogger) *AuthService {
	return &AuthService{
		Repo:   r,
		logger: l,
	}
}

func (service *AuthService) GenerateHashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword), err
}

func (service *AuthService) CheckHashPassword(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err != nil
}

func (service *AuthService) GenerateTokens(username string, userId string) (accessToken string, refreshToken string, err error) {
	accessToken, err = utils.GenerateJWT(username, userId, false)
	if err != nil {
		return "", "", err
	}
	refreshToken, err = utils.GenerateJWT(username, userId, true)
	if err != nil {
		return "", "", err
	}
	return
}

func (service *AuthService) Create(ctx *gin.Context) {
	var input dto.CreateUserRequest

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
	if input.Password != input.PasswordRepeat {
		ctx.JSON(http.StatusBadRequest, "Passwords do not match")
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

func (service *AuthService) Authorize(ctx *gin.Context) {
	var input dto.SignInRequest
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
	user, err := service.Repo.GetForceByName(input.Username)
	if err != nil {
		service.logger.Errorf("Couldn't find user with username %s", input.Username)
		ctx.Status(http.StatusNotFound)
		return
	}
	if service.CheckHashPassword(input.Password, user.Password) {
		ctx.JSON(http.StatusBadRequest, "Invalid username or/and password")
		return
	}
	accessToken, refreshToken, err := service.GenerateTokens(user.Username, user.UUID.String())
	if err != nil {
		service.logger.Errorf("Couldn't generate tokens due to %s", err.Error())
		ctx.Status(http.StatusInternalServerError)
		return
	}
	ctx.SetCookie(
		"refreshToken",
		refreshToken,
		3600,
		"/",
		ctx.ClientIP(),
		false,
		true,
	)
	ctx.JSON(200, dto.SignInResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
}

func (service *AuthService) RefreshToken(ctx *gin.Context) {
	refreshTokenFromCookie, err := ctx.Cookie("refreshToken")
	if err != nil {
		ctx.JSON(400, "Couldn't extract refresh token")
		return
	}
	claims, validationErr := utils.ValidateToken(refreshTokenFromCookie)
	if validationErr != nil {
		ctx.JSON(400, "Refresh token is not valid")
		return
	}
	accessToken, refreshToken, tokensErr := service.GenerateTokens(claims.Username, claims.UserId)
	if tokensErr != nil {
		service.logger.Errorf("Couldn't generate tokens due to %s", err.Error())
		ctx.Status(http.StatusInternalServerError)
		return
	}
	ctx.SetCookie(
		"refreshToken",
		refreshToken,
		3600,
		"/",
		ctx.ClientIP(),
		false,
		true,
	)
	ctx.JSON(200, dto.SignInResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
}
