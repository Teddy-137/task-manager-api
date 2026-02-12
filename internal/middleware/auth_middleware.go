package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/teddy-137/task_manager_api/pkg/utils"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Missing authentication token"})
			return
		}

		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authentication token"})
			ctx.Abort()
			return
		}

		tokenString := tokenParts[1]

		claims, err := utils.VerifyToken(tokenString)

		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authentication token"})
			ctx.Abort()
			return
		}

		ctx.Set("user_id", claims["user_id"])

		ctx.Next()
	}

}
