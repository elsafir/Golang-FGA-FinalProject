package middleware

import (
	"Golang-FGA-FinalProject/helpers"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("Authorization")
		if token == "" {
			helpers.ResponseStatusUnauthorized(ctx, "Token Not Found")
			return
		}

		tokenStr := strings.Split(token, "Bearer ")[1]
		if tokenStr == "" {
			helpers.ResponseStatusUnauthorized(ctx, "Token Not Found")
			return
		}

		claims, err := helpers.VerifyToken(tokenStr)
		if err != nil {
			helpers.ResponseStatusUnauthorized(ctx, err.Error())
		}

		var data = claims.(jwt.MapClaims)
		ctx.Set("id", data["id"])
		ctx.Set("email", data["email"])
		ctx.Next()
	}
}
