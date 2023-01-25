package errors

import (
	"github.com/gin-gonic/gin"
	"github.com/insmnia/go-users-service/core/validators"
	"net/http"
)

type ParseBodyError struct {
	Msg string
}

func (pbe *ParseBodyError) Raise(c *gin.Context) {
	c.JSON(http.StatusBadRequest, pbe.Msg)
}

type ValidateBodyError struct {
	Errors []*validators.ErrorResponse
}

func (vbe *ValidateBodyError) Raise(c *gin.Context) {
	c.JSON(http.StatusUnprocessableEntity, vbe.Errors)
}
