package middleware

import (
	"api/modules/jwt"
	"api/modules/response"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")

		if authHeader == "" {
			response.Error(ctx, 401, "Invalid Token", nil)
			ctx.Abort()
			return 
		}

		tokenString := strings.Split(authHeader, " ")
		if len(tokenString) != 2 {
			response.Error(ctx, 401, "Invalid Token", nil)
			ctx.Abort()
			return 
		}

		claims, err := jwt.ValidateToken(tokenString[1])
		if err != nil {
			response.Error(ctx, 401, "Invalid Token", nil)
			ctx.Abort()
			return 
		}

		ctx.Set("userId", claims.UserID)
		ctx.Set("email", claims.Email)
	}
}