package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/insmnia/go-users-service/pkg/utils"
	"strings"
)

func RequiresAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")
		if tokenString == "" {
			ctx.Status(401)
			ctx.Abort()
			return
		}
		tokenParts := strings.Split(tokenString, " ")
		claims, err := utils.ValidateToken(tokenParts[1])
		if err != nil {
			ctx.JSON(401, gin.H{"error": err.Error()})
			ctx.Abort()
			return
		}
		ctx.Set("user", claims.UserId)
		ctx.Next()
	}
}
